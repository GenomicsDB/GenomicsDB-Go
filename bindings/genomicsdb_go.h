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

#pragma once

#include <stdint.h>
#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

  const char *genomicsdb_version();

#define PATH_MAX 4095
  typedef struct genomicsdb_status_t{
    int succeeded;
    char error_message[PATH_MAX+1];
  } genomicsdb_status_t;

  typedef struct genomicsdb_info_t{
    int kind; //strings=0/ints=1/floats=2
    char *name;
    void *ptr;
  } genomicsdb_info_t;

  void *genomicsdb_connect(void *pb_string, size_t len, genomicsdb_status_t *status);

  void *genomicsdb_query(void *genomicsdb_handle, genomicsdb_status_t *status);

  uint64_t genomicsdb_get_count(void *query_processor);

  char *genomicsdb_get_sample_name_at(void *query_processor, uint64_t index);
  char *genomicsdb_get_chromosome_at(void *query_processor, uint64_t index);
  int64_t *genomicsdb_get_positions(void *query_processor);
  int64_t *genomicsdb_get_end_positions(void *query_processor);

  uint64_t genomicsdb_get_genomic_field_count(void *query_processor);
  int genomicsdb_get_genomic_field_info(void *query_processor, uint64_t index, genomicsdb_info_t *info_t);
  char *genomicsdb_get_genomic_string_field_at(void *query_processor, char *field_name, uint64_t index);
  
  void genomicsdb_delete_query(void *query_processor);

  void genomicsdb_disconnect(void *genomicsdb);

#ifdef __cplusplus
}
#endif
