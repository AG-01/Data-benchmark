import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface ResultState {
  results: any[];
}

const initialState: ResultState = {
  results: [],
};

const resultSlice = createSlice({
  name: 'results',
  initialState,
  reducers: {
    setResults: (state, action: PayloadAction<any[]>) => {
      state.results = action.payload;
    },
  },
});

export const { setResults } = resultSlice.actions;
export default resultSlice.reducer;
