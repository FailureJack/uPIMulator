/***************************************************************************
 *cr
 *cr            (C) Copyright 2015 The Board of Trustees of the
 *cr                        University of Illinois
 *cr                         All Rights Reserved
 *cr
 ***************************************************************************/
/*
  In-Place Data Sliding Algorithms for Many-Core Architectures, presented in ICPP’15

  Copyright (c) 2015 University of Illinois at Urbana-Champaign. 
  All rights reserved.

  Permission to use, copy, modify and distribute this software and its documentation for 
  educational purpose is hereby granted without fee, provided that the above copyright 
  notice and this permission notice appear in all copies of this software and that you do 
  not sell the software.

  THE SOFTWARE IS PROVIDED "AS IS" AND WITHOUT WARRANTY OF ANY KIND,EXPRESS, IMPLIED OR 
  OTHERWISE.

  Authors: Juan Gómez-Luna (el1goluj@uco.es, gomezlun@illinois.edu), Li-Wen Chang (lchang20@illinois.edu)
*/

__global__ void select_remove_if(T *matrix_out, T *matrix,
    int size,
    volatile unsigned int *flags,
    struct is_even pred)
{
  __shared__ int count; // Counter for number of non-zero elements per block
  const int num_flags = size % (blockDim.x * REGS) == 0 ? size / (blockDim.x * REGS) : size / (blockDim.x * REGS) + 1;

  // Dynamic allocation of runtime workgroup id
  if (threadIdx.x == 0) count = 0;
  const int my_s = dynamic_wg_id(flags, num_flags);

  int local_cnt = 0;
  // Declare on-chip memory
  T reg[REGS];
  int pos = my_s * REGS * blockDim.x + threadIdx.x;
  // Load in on-chip memory
  #pragma unroll
  for (int j = 0; j < REGS; j++){
    if (pos < size){
      reg[j] = matrix[pos];
      if(!pred(reg[j]))
        local_cnt++;
      else
        reg[j] = -1;
    }
    else
      reg[j] = -1;
    pos += blockDim.x;
  }
  reduction<int>(&count, local_cnt);

  // Set global synch
  ds_sync_irregular(flags, my_s, &count);

  // Store to global memory 
  #pragma unroll
  for (int j = 0; j < REGS; j++){
    pos = block_binary_prefix_sums(&count, reg[j] >= 0);
    if (reg[j] >= 0){
      matrix_out[pos] = reg[j];
    }
  }
}

__global__ void select_copy_if(T *matrix_out, T *matrix,
    int size,
    volatile unsigned int *flags,
    struct is_even pred)
{
  __shared__ int count; // Counter for number of non-zero elements per block
  const int num_flags = size % (blockDim.x * REGS) == 0 ? size / (blockDim.x * REGS) : size / (blockDim.x * REGS) + 1;

  // Dynamic allocation of runtime workgroup id
  if (threadIdx.x == 0) count = 0;
  const int my_s = dynamic_wg_id(flags, num_flags);

  int local_cnt = 0;
  // Declare on-chip memory
  T reg[REGS];
  int pos = my_s * REGS * blockDim.x + threadIdx.x;
  // Load in on-chip memory
  #pragma unroll
  for (int j = 0; j < REGS; j++){
    if (pos < size){
      reg[j] = matrix[pos];
      if(pred(reg[j]))
        local_cnt++;
      else
        reg[j] = -1;
    }
    else
      reg[j] = -1;
    pos += blockDim.x;
  }
  reduction<int>(&count, local_cnt);

  // Set global synch
  ds_sync_irregular(flags, my_s, &count);

  // Store to global memory 
  #pragma unroll
  for (int j = 0; j < REGS; j++){
    pos = block_binary_prefix_sums(&count, reg[j] >= 0);
    if (reg[j] >= 0){
      matrix_out[pos] = reg[j];
    }
  }
}
