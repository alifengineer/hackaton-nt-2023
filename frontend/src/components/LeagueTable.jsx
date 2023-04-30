import React, { useEffect, useState } from "react";
import MenuItem from "@mui/material/MenuItem";
import FormControl from "@mui/material/FormControl";
import Select from "@mui/material/Select";
import defImg from "../images/default_user.png";

const LeagueTable = ({ leagues, teams, handleTable }) => {
  const [age, setAge] = useState("");

  useEffect(() => {
    setAge(leagues[0]?.id);
  }, [leagues]);

  const handleChange = (event) => {
    setAge(event.target.value);
    handleTable(event.target.value);
  };

  return (
    <div className="league__table">
      <h2 className="table__title">Jadval</h2>

      <FormControl sx={{ minWidth: 120 }} fullWidth>
        <Select
          value={age}
          onChange={handleChange}
          displayEmpty
          sx={{
            fontSize: "16px",
            backgroundColor: "#fff",
            height: "40px !important",
          }}
        >
          {leagues?.map((val, index) => {
            return (
              <MenuItem key={index} sx={{ fontSize: "14px" }} value={val?.id}>
                {val.name}
              </MenuItem>
            );
          })}
        </Select>
      </FormControl>
      <div className="table-container">
        <table className="table">
          <thead style={{ height: "40px" }}>
            <tr>
              <th style={{ width: "25px" }}>№</th>
              <th></th>
              <th>Команда</th>
              <th>И</th>
              <th>О</th>
            </tr>
          </thead>
          <tbody>
            {teams.map((row, index) => (
              <tr
                key={index + 1}
                className={`${index % 2 === 0 && "odd__row"}`}
              >
                <td style={{ width: "25px" }}>{index + 1}</td>
                <td style={{ width: "40px" }}>
                  <img
                    style={{
                      width: "30px",
                      height: "30px",
                      borderRadius: "50%",
                    }}
                    src={row.image || defImg}
                    alt=""
                  />
                </td>
                <td>{row.name}</td>
                <td style={{ width: "20px" }}>{row.played}</td>
                <td style={{ width: "20px" }}>{row.score}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default LeagueTable;
