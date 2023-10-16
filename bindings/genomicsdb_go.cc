/**
 *
 * The MIT License
 *
 * Copyright (c) 2023 dātma, inc™
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 *
 */

#include "genomicsdb_go.h"

#include "genomicsdb.h"

#include <cmath>
#include <stdio.h>
#include <string.h>
#include <time.h>

class CountCellsProcessor : public GenomicsDBVariantCallProcessor {
 public:
  CountCellsProcessor() {
  };

  void process(const interval_t& interval) {
    m_intervals++;
  };

  void process(const std::string& sample_name,
               const int64_t* coordinates,
               const genomic_interval_t& genomic_interval,
               const std::vector<genomic_field_t>& genomic_fields) {
    m_count++;
  };

  int m_intervals = 0;
  int m_count = 0;
};

#define STRING_FIELD(NAME, TYPE) (TYPE.is_string() || TYPE.is_char() || TYPE.num_elements > 1 || (NAME.compare("GT") == 0))
#define INT_FIELD(TYPE) (TYPE.is_int())
#define FLOAT_FIELD(TYPE) (TYPE.is_float())

class VariantCallProcessor : public GenomicsDBVariantCallProcessor {
  void process(const interval_t& interval) {
    if (!m_is_initialized) {
      m_is_initialized = true;
      auto& genomic_field_types = get_genomic_field_types();
      for (auto& field_type_pair : *genomic_field_types) {
        std::string field_name = field_type_pair.first;
        genomic_field_type_t field_type = field_type_pair.second;
        if (field_name.compare("END")==0) {
          continue;
        }
        // Order fields by inserting REF and ALT in the beginning
        if (!field_name.compare("REF") && m_field_names.size() > 1) {
          m_field_names.insert(m_field_names.begin(), field_name);
        } else if (!field_name.compare("ALT") && m_field_names.size() > 2) {
          m_field_names.insert(m_field_names.begin()+1, field_name);
        } else {
          m_field_names.push_back(field_name);
        }
        if (STRING_FIELD(field_name, field_type)) {
          std::vector<std::string> str_vector;
          m_string_fields.emplace(std::make_pair(field_name, std::move(str_vector))) ;
        } else if (INT_FIELD(field_type)) {
          std::vector<int64_t> int_vector;
          m_int_fields.emplace(std::make_pair(field_name, std::move(int_vector)));
        } else if (FLOAT_FIELD(field_type)) {
          std::vector<float> float_vector;
          m_float_fields.emplace(std::make_pair(field_name, std::move(float_vector)));
        } else {
          std::string msg = "Genomic field type for " + field_name + " not supported";
          throw std::runtime_error(msg.c_str());
        }
      }
    }
  }

  void process_fields(const std::vector<genomic_field_t>& genomic_fields) {
    for (auto field_name: m_field_names) {
      // END is part of the Genomic Coordinates, so don't process here
      if (field_name.compare("END") == 0) {
        continue;
      }
        
      auto field_type = get_genomic_field_types()->at(field_name);
      
      bool found = false;
      for (auto genomic_field: genomic_fields) {
        if (genomic_field.name.compare(field_name) == 0) {
          if (STRING_FIELD(field_name, field_type)) {
            m_string_fields[field_name].push_back(genomic_field.to_string(field_type).c_str());
          } else if (INT_FIELD(field_type)) {
            m_int_fields[field_name].push_back(genomic_field.int_value_at(0));
          } else if (FLOAT_FIELD(field_type)) {
            m_float_fields[field_name].push_back( genomic_field.float_value_at(0));
          } else {
            std::string msg = "Genomic field type for " + field_name + " not supported";
            throw std::runtime_error(msg.c_str());
          }
          found = true;
          break;
        }
      }
      
      if (!found) {
        if (STRING_FIELD(field_name, field_type)) {
          m_string_fields[field_name].push_back("");
        } else if (INT_FIELD(field_type)) {
          m_int_fields[field_name].push_back(-99999);
        } else if (FLOAT_FIELD(field_type)) {
          m_float_fields[field_name].push_back(std::nanf(""));
        } else {
          std::string msg = "Genomic field type for " + field_name + " not supported";
          throw std::runtime_error(msg.c_str());
        }
      }
    }
  }
  
  void process(const std::string& sample_name,
               const int64_t* coordinates,
               const genomic_interval_t& genomic_interval,
               const std::vector<genomic_field_t>& genomic_fields) {
    m_rows.push_back(coordinates[0]);
    m_cols.push_back(coordinates[1]);
    m_sample_names.push_back(sample_name);
    m_chrom.push_back(genomic_interval.contig_name);
    m_pos.push_back(genomic_interval.interval.first);
    m_end.push_back(genomic_interval.interval.second);
    
    process_fields(genomic_fields);
  }

 public:
  size_t count() {
    return m_rows.size();
  }

  std::vector<int64_t> m_rows;
  std::vector<int64_t> m_cols;
  std::vector<std::string> m_sample_names;
  std::vector<std::string> m_chrom;
  std::vector<int64_t> m_pos;
  std::vector<int64_t> m_end;
  std::vector<std::string> m_field_names;
  std::map<std::string, std::vector<std::string>> m_string_fields;
  std::map<std::string, std::vector<int64_t>> m_int_fields;
  std::map<std::string, std::vector<float>> m_float_fields;

 private:
  bool m_is_initialized = false;
};

const char *genomicsdb_version() {
  static std::string version = genomicsdb::version();
  return version.c_str();
}

void *genomicsdb_connect(void *pb_string, size_t len, genomicsdb_status_t *status) {
  time_t start = time(0);
  void *genomicsdb = 0;
  try {
    status->succeeded = 1;
    genomicsdb = new GenomicsDB(std::string((char *)pb_string, len), GenomicsDB::PROTOBUF_BINARY_STRING);
    CountCellsProcessor count_cells_processor;
  } catch (const std::exception &e) {
    status->succeeded = 0;
    strncpy(&status->error_message[0], e.what(), PATH_MAX);
  }
  printf("Connect time=%lus\n", time(0)-start);
  return genomicsdb;
}

void *genomicsdb_query(void *genomicsdb_handle, genomicsdb_status_t *status) {
  VariantCallProcessor *variant_call_processor = 0;
  try {
    status->succeeded = 1;
    variant_call_processor = new VariantCallProcessor();
    ((GenomicsDB *)genomicsdb_handle)->query_variant_calls(*variant_call_processor);
  } catch (const std::exception &e) {
    status->succeeded = 0;
    strncpy(&status->error_message[0], e.what(), PATH_MAX);
  }
  return variant_call_processor;
}

#define VARIANT_CALL_PROCESSOR ((VariantCallProcessor *)query_processor)

uint64_t genomicsdb_get_count(void *query_processor) {
  return ((VariantCallProcessor *)query_processor)->count();
}

char *genomicsdb_get_sample_name_at(void *query_processor, uint64_t index) {
  if (index < VARIANT_CALL_PROCESSOR->m_sample_names.size()) {
    auto& sample_names = VARIANT_CALL_PROCESSOR->m_sample_names;
    return const_cast<char *>(sample_names[index].c_str());
  }
  return 0;
}

char *genomicsdb_get_chromosome_at(void *query_processor, uint64_t index) {
  if (index < VARIANT_CALL_PROCESSOR->m_sample_names.size()) {
    auto& chrom = VARIANT_CALL_PROCESSOR->m_chrom;
    return const_cast<char *>(chrom[index].c_str());
  }
  return 0;
}

int64_t *genomicsdb_get_positions(void *query_processor) {
  auto& pos = VARIANT_CALL_PROCESSOR->m_pos;
  return &pos[0];
}

int64_t *genomicsdb_get_end_positions(void *query_processor) {
  auto& end = VARIANT_CALL_PROCESSOR->m_end;
  return &end[0];
}

uint64_t genomicsdb_get_genomic_field_count(void *query_processor) {
  return VARIANT_CALL_PROCESSOR->m_field_names.size();
}

int genomicsdb_get_genomic_field_info(void *query_processor, uint64_t index, genomicsdb_info_t *info) {
  int found = 1;
  auto& field_name = VARIANT_CALL_PROCESSOR->m_field_names[index];
  info->name = &field_name[0];
  if (VARIANT_CALL_PROCESSOR->m_string_fields.find(field_name) != VARIANT_CALL_PROCESSOR->m_string_fields.end()) {
    info->kind = 0;
    info->ptr = 0; // Use get_genomic_string_field_at() to get the string based on the index
  } else if (VARIANT_CALL_PROCESSOR->m_int_fields.find(field_name) != VARIANT_CALL_PROCESSOR->m_int_fields.end()) {
    auto& int_field = VARIANT_CALL_PROCESSOR->m_int_fields[field_name];
    info->kind = 1;
    info->ptr = &int_field[0];
  } else  if (VARIANT_CALL_PROCESSOR->m_float_fields.find(field_name) != VARIANT_CALL_PROCESSOR->m_float_fields.end()) {
    auto& float_field = VARIANT_CALL_PROCESSOR->m_float_fields[field_name];
    info->kind = 2;
    info->ptr = &float_field[0];
  } else {
    found = 0;
  }
  return found;
}

char *genomicsdb_get_genomic_string_field_at(void *query_processor, char *field_name, uint64_t index) {
  if (index <  VARIANT_CALL_PROCESSOR->m_string_fields[field_name].size()) {
    auto& string_field = VARIANT_CALL_PROCESSOR->m_string_fields[field_name];
    return const_cast<char *>(string_field[index].c_str());
  }
  return 0;
}

void genomicsdb_delete_query(void *query_processor) {
  delete (VariantCallProcessor *)query_processor;
}

void genomicsdb_disconnect(void *genomicsdb) {
  delete (GenomicsDB *)genomicsdb;
}
