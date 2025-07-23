import React from 'react';
import { Container, Typography } from '@mui/material';

const Results: React.FC = () => {
  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Typography variant="h4" component="h1" gutterBottom>
        Results
      </Typography>
      {/* Placeholder for results display */}
    </Container>
  );
};

export default Results;
