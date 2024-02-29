#ifndef UPMEM_SIM_ABI_ISA_EXCEPTION_H_
#define UPMEM_SIM_ABI_ISA_EXCEPTION_H_

namespace upmem_sim::abi::isa {

enum Exception {
  MEMORY_FAULT,
  DMA_FAULT,
  HEAP_FULL,
  DIVISION_BY_ZERO,
  ASSERT,
  HALT,
  PRINT_OVERFLOW,
  ALREADY_PROFILING,
  NOT_PROFILING
};
}

#endif
