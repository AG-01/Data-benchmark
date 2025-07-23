import React from 'react';
import { Container, Typography } from '@mui/material';

const Benchmarks: React.FC = () => {
  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Typography variant="h4" component="h1" gutterBottom>
        Benchmarks
      </Typography>
      {/* Placeholder for benchmarks list */}
    </Container>
  );
};

export default Benchmarks;
