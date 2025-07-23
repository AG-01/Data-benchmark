import React from 'react';
import { Container, Typography } from '@mui/material';

const QueryEditor: React.FC = () => {
  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Typography variant="h4" component="h1" gutterBottom>
        Query Editor
      </Typography>
      {/* Placeholder for query editor */}
    </Container>
  );
};

export default QueryEditor;
