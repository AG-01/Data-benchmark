import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface QueryState {
  queries: any[];
  selectedQuery: any | null;
}

const initialState: QueryState = {
  queries: [],
  selectedQuery: null,
};

const querySlice = createSlice({
  name: 'queries',
  initialState,
  reducers: {
    setQueries: (state, action: PayloadAction<any[]>) => {
      state.queries = action.payload;
    },
    setSelectedQuery: (state, action: PayloadAction<any | null>) => {
      state.selectedQuery = action.payload;
    },
  },
});

export const { setQueries, setSelectedQuery } = querySlice.actions;
export default querySlice.reducer;
