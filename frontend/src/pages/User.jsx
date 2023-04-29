import React from 'react';
import Button from '@mui/material/Button';
import { Header, Footer } from '../components';

const User = () => {
  return (
    <div>
      <h2>User</h2>
      <Footer />
      <Header />
      <Button variant="contained">Hello World</Button>
    </div>
  );
}

export default User;