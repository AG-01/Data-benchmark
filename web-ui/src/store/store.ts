import { configureStore } from '@reduxjs/toolkit';
import benchmarkReducer from './slices/benchmarkSlice';
import queryReducer from './slices/querySlice';
import resultReducer from './slices/resultSlice';
import engineReducer from './slices/engineSlice';
import uiReducer from './slices/uiSlice';

export const store = configureStore({
  reducer: {
    benchmarks: benchmarkReducer,
    queries: queryReducer,
    results: resultReducer,
    engines: engineReducer,
    ui: uiReducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: {
        ignoredActions: ['persist/PERSIST'],
      },
    }),
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
