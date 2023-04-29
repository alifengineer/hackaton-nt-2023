import { useState } from "react";
import { useEffect } from "react";
import { createContext } from "react";
import { authAxios } from "../utils/axios";

const AuthContext = createContext({});

export const AuthProvider = ({ children }) => {
  const [admin, setAdmin] = useState(null);

  useEffect(() => {

  }, [])
}