import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { Benchmark, BenchmarkStatus } from '../../types';
import api from '../../services/api';

interface BenchmarkState {
  benchmarks: Benchmark[];
  currentBenchmark: Benchmark | null;
  status: BenchmarkStatus | null;
  loading: boolean;
  error: string | null;
}

const initialState: BenchmarkState = {
  benchmarks: [],
  currentBenchmark: null,
  status: null,
  loading: false,
  error: null,
};

// Async thunks
export const fetchBenchmarks = createAsyncThunk(
  'benchmarks/fetchBenchmarks',
  async (params?: { status?: string; table_format?: string; limit?: number; offset?: number }) => {
    const response = await api.get('/benchmarks', { params });
    return response.data;
  }
);

export const fetchBenchmarkById = createAsyncThunk(
  'benchmarks/fetchBenchmarkById',
  async (id: number) => {
    const response = await api.get(`/benchmarks/${id}`);
    return response.data;
  }
);

export const createBenchmark = createAsyncThunk(
  'benchmarks/createBenchmark',
  async (benchmark: Omit<Benchmark, 'id' | 'created_at' | 'updated_at'>) => {
    const response = await api.post('/benchmarks', benchmark);
    return response.data;
  }
);

export const updateBenchmark = createAsyncThunk(
  'benchmarks/updateBenchmark',
  async (benchmark: Benchmark) => {
    const response = await api.put(`/benchmarks/${benchmark.id}`, benchmark);
    return response.data;
  }
);

export const deleteBenchmark = createAsyncThunk(
  'benchmarks/deleteBenchmark',
  async (id: number) => {
    await api.delete(`/benchmarks/${id}`);
    return id;
  }
);

export const runBenchmark = createAsyncThunk(
  'benchmarks/runBenchmark',
  async (id: number) => {
    const response = await api.post(`/benchmarks/${id}/run`);
    return response.data;
  }
);

export const fetchBenchmarkStatus = createAsyncThunk(
  'benchmarks/fetchBenchmarkStatus',
  async (id: number) => {
    const response = await api.get(`/benchmarks/${id}/status`);
    return response.data;
  }
);

const benchmarkSlice = createSlice({
  name: 'benchmarks',
  initialState,
  reducers: {
    clearError: (state) => {
      state.error = null;
    },
    setCurrentBenchmark: (state, action: PayloadAction<Benchmark | null>) => {
      state.currentBenchmark = action.payload;
    },
  },
  extraReducers: (builder) => {
    builder
      // Fetch benchmarks
      .addCase(fetchBenchmarks.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(fetchBenchmarks.fulfilled, (state, action: PayloadAction<Benchmark[]>) => {
        state.loading = false;
        state.benchmarks = action.payload;
      })
      .addCase(fetchBenchmarks.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error.message || 'Failed to fetch benchmarks';
      })
      // Fetch benchmark by ID
      .addCase(fetchBenchmarkById.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(fetchBenchmarkById.fulfilled, (state, action: PayloadAction<Benchmark>) => {
        state.loading = false;
        state.currentBenchmark = action.payload;
      })
      .addCase(fetchBenchmarkById.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error.message || 'Failed to fetch benchmark';
      })
      // Create benchmark
      .addCase(createBenchmark.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(createBenchmark.fulfilled, (state, action: PayloadAction<Benchmark>) => {
        state.loading = false;
        state.benchmarks.push(action.payload);
      })
      .addCase(createBenchmark.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error.message || 'Failed to create benchmark';
      })
      // Update benchmark
      .addCase(updateBenchmark.fulfilled, (state, action: PayloadAction<Benchmark>) => {
        const index = state.benchmarks.findIndex(b => b.id === action.payload.id);
        if (index !== -1) {
          state.benchmarks[index] = action.payload;
        }
        if (state.currentBenchmark?.id === action.payload.id) {
          state.currentBenchmark = action.payload;
        }
      })
      // Delete benchmark
      .addCase(deleteBenchmark.fulfilled, (state, action: PayloadAction<number>) => {
        state.benchmarks = state.benchmarks.filter(b => b.id !== action.payload);
        if (state.currentBenchmark?.id === action.payload) {
          state.currentBenchmark = null;
        }
      })
      // Run benchmark
      .addCase(runBenchmark.fulfilled, (state) => {
        // Update status or trigger status fetch
      })
      // Fetch benchmark status
      .addCase(fetchBenchmarkStatus.fulfilled, (state, action: PayloadAction<BenchmarkStatus>) => {
        state.status = action.payload;
      });
  },
});

export const { clearError, setCurrentBenchmark } = benchmarkSlice.actions;
export default benchmarkSlice.reducer;
