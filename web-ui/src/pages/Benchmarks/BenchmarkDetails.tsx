import React from 'react';
import { Container, Typography } from '@mui/material';

const BenchmarkDetails: React.FC = () => {
  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Typography variant="h4" component="h1" gutterBottom>
        Benchmark Details
      </Typography>
      {/* Placeholder for benchmark details */}
    </Container>
  );
};

export default BenchmarkDetails;
