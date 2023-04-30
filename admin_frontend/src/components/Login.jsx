import React, { useState } from "react";
import { useAuth } from "../context/auth";

const Login = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = (event) => {
    event.preventDefault();
    // Add login logic here
  };

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

      <button type="submit" className="btn btn-primary mt-3">
        Login
      </button>
    </form>
  );
};

export default Login;
