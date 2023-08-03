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

#include <stdio.h>
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

const char *version() {
  static std::string version = genomicsdb::version();
  return version.c_str();
}

void *connect(void *pb_string, size_t len) {
  time_t start = time(0);
  auto genomicsdb = new GenomicsDB(std::string((char *)pb_string, len), GenomicsDB::PROTOBUF_BINARY_STRING);
  CountCellsProcessor count_cells_processor;
  printf("Connect time=%lus\n", time(0)-start);
  return genomicsdb;
}

void query_variant_calls(void* genomicsdb_handle) {
  time_t start = time(0);
  CountCellsProcessor count_cells_processor;
  ((GenomicsDB *)genomicsdb_handle)->query_variant_calls(count_cells_processor);
  printf("Query Variant Calls time=%lus\n", time(0)-start);
  printf("Count of cells=%d\n", count_cells_processor.m_count);
}

void disconnect(void *genomicsdb) {
  delete (GenomicsDB *)genomicsdb;
}
