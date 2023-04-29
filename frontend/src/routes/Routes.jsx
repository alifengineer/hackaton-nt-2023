import React from 'react';
import { Routes, Route } from 'react-router-dom';
import Admin from '../components/Admin/Admin';
import User from '../components/User/User';
export default function Routers() {
  return (
    <Routes>
      <Route index element={<User />} />
      <Route path='/admin' element={<Admin />} />
    </Routes>
  )
}
