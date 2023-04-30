import { createContext, useState, useContext, useEffect } from "react";
import Cookies from "js-cookie";
import { authAxios } from "../utils/axios";
import toast from "react-hot-toast";
import { useLocation } from "react-router-dom";
import { useNavigate } from "react-router-dom";

const AuthContext = createContext({});

export const AuthProvider = ({ children }) => {
  const [user, setUser] = useState(false);
  const [loading, setLoading] = useState();
  const [formError, setFormError] = useState("");
  const [redirect, setRedirect] = useState("/");

  // useEffect(() => {
  //   async function loadUserFromCookies() {
  //     const token = Cookies.get("token");
  //     if (token) {
  //       authAxios.defaults.headers.Authorization = `Bearer ${token}`;
  //       try {
  //         setUser(true);
  //       } catch (err) {
  //         setUser(false);
  //         Cookies.remove("token");
  //       }
  //     } else {
  //       setUser(null);
  //     }
  //     setLoading(false);
  //   }

  //   loadUserFromCookies();
  // }, []);

  const login = async () => {};

  return (
    <AuthContext.Provider
      value={{
        user,
        setUser,
        loading,
        setLoading,
        setRedirect,
        redirect,
        login,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => useContext(AuthContext);

export const AuthGuard = ({ children }) => {
  const { user, loading, setRedirect } = useAuth();
  const location = useLocation();
  const navigate = useNavigate();

  useEffect(() => {
    console.log(location);
    if (!loading) {
      if (!user) {
        setRedirect(location ? location.pathname : "/");
        navigate("/login");
      }
    }
  }, [loading, location, user, setRedirect, navigate]);

  if (loading) {
    return <div></div>;
  }

  if (!loading && user) {
    return <>{children}</>;
  }

  return null;
};
