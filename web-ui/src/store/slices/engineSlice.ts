import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface EngineState {
  engines: any[];
}

const initialState: EngineState = {
  engines: [],
};

const engineSlice = createSlice({
  name: 'engines',
  initialState,
  reducers: {
    setEngines: (state, action: PayloadAction<any[]>) => {
      state.engines = action.payload;
    },
  },
});

export const { setEngines } = engineSlice.actions;
export default engineSlice.reducer;
