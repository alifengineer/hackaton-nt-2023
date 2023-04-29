import React from 'react';
import { Routes, Route } from 'react-router-dom';
import Admin from '../pages/Admin';
import User from '../pages/User';
export default function Routers() {
  return (
    <Routes>
      <Route index element={<User />} />
      <Route path='/admin' element={<Admin />} />
    </Routes>
  )
}
