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

  const char *version();

#define PATH_MAX 4095
  typedef struct status_t{
    int succeeded;
    char error_message[PATH_MAX+1];
  } status_t;

  typedef struct info_t{
    int kind; //strings=0/ints=1/floats=2
    char *name;
    void *ptr;
  } info_t;

  void *connect(void *pb_string, size_t len, status_t *status);

  void *query(void *genomicsdb_handle, status_t *status);

  uint64_t get_count(void *query_processor);

  char *get_sample_name_at(void *query_processor, uint64_t index);
  char *get_chromosome_at(void *query_processor, uint64_t index);
  int64_t *get_positions(void *query_processor);
  int64_t *get_end_positions(void *query_processor);

  uint64_t get_genomic_field_count(void *query_processor);
  int get_genomic_field_info(void *query_processor, uint64_t index, info_t *info_t);
  char *get_genomic_string_field_at(void *query_processor, char *field_name, uint64_t index);
  
  void delete_query(void *query_processor);

  void disconnect(void *genomicsdb);

#ifdef __cplusplus
}
#endif
