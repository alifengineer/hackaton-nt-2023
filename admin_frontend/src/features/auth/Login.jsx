import React, { useState } from "react";
import axios, {authAxios, simpleAxios} from "../../utils/axios";
import Cookies from "js-cookie";
import toast from "react-hot-toast";
import {useNavigate} from "react-router-dom";

const Login = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [smsId, setSmsId] = useState(null);
  const [OTP, setOTP] = useState("");
  const [step, setStep] = useState(1);
  const navigate = useNavigate();

  const handleSubmit = async (event) => {
    event.preventDefault();
    try {
        await authAxios.post(`api/v1/email/verify-email/${smsId}/${OTP}`)
        const {data: {data: {access_token}}} = await authAxios.post("api/v1/auth/login", {
            email: email,
            password: password
        })

        Cookies.set("token", access_token);
        simpleAxios.interceptors.request.use((config) => {
            config.headers.Authorization = `Bearer ${access_token}`
            return config;
        })
        axios.interceptors.request.use((config) => {
            config.headers.Authorization = `Bearer ${access_token}`
            return config;
        })
        authAxios.interceptors.request.use((config) => {
            config.headers.Authorization = `Bearer ${access_token}`
            return config;
        })

        navigate("/");
        toast.success("Успешно авторизован");
    } catch (error) {
        toast.error("Ошибка при авторизации")
        console.log(error);
    }
  };

  const updateStep = async () => {
      try {
          const {data: {data: {sms_id}}} = await authAxios.post("api/v1/email/send-otp", {
              email: email
          });
          setSmsId(sms_id);
          setStep(2);
      } catch (error) {
          console.log(error);
      }
  }

    return (
        <form
            onSubmit={handleSubmit}
            style={{
                width: "500px",
                padding: "20px",
                border: "1px solid #333",
                margin: "0 auto",
                marginTop: "50px",
                borderRadius: "15px",
            }}
        >
            <h4 className="login" style={{ textAlign: "center" }}>
                Login
            </h4>
            {step === 1 && <>
                <div className="form-group mt-3">
                    <label htmlFor="email">Email address</label>
                    <input
                        type="email"
                        className="form-control"
                        id="email"
                        placeholder="Enter email"
                        value={email}
                        required={true}
                        onChange={(event) => setEmail(event.target.value)}
                    />
                </div>

                <div className="form-group mt-3">
                    <label htmlFor="password">Password</label>
                    <input
                        type="password"
                        className="form-control"
                        id="password"
                        placeholder="Password"
                        value={password}
                        required={true}
                        onChange={(event) => setPassword(event.target.value)}
                    />
                </div>

                <button onClick={updateStep} type="button" className="btn btn-primary mt-3">
                    Login
                </button>
            </>}
            {step === 2 && <>
                <div className="form-group mt-3">
                    <label htmlFor="email">OTP</label>
                    <input
                        type="text"
                        className="form-control"
                        id="otp"
                        placeholder="Enter otp"
                        value={OTP}
                        required={true}
                        onChange={(event) => setOTP(event.target.value)}
                    />
                </div>
                <button type="submit" className="btn btn-primary mt-3">
                    Login
                </button>
            </>}
        </form>
    );
};

export default Login;
