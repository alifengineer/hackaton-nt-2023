import React, { useEffect, useState } from "react";
import { Button, Col, Form, Row } from "react-bootstrap";
import { SelectLeague, SelectTeam } from "../../match/components";
import axios from "../../../utils/axios";
import { Controller, useFormContext, useWatch } from "react-hook-form";

const TeamForm = () => {
  const [leagues, setLeagues] = useState([]);
  const [teams, setTeams] = useState([]);

  const leagueId = useWatch({ name: "leagueId" });
  const form = useFormContext();

  useEffect(() => {
    const getAllTeams = async () => {
      const {
        data: {
          data: { teams },
        },
      } = await axios.get("api/v1/team/");
      setTeams(teams);
    };

    const getAllLeagues = async () => {
      const {
        data: { data },
      } = await axios.get("api/v1/league/");
      setLeagues(data);
    };

    getAllTeams();
    getAllLeagues();
  }, []);

  useEffect(() => {
    const getLeague = async () => {
      const {
        data: { data },
      } = await axios.get(`api/v1/league/${leagueId}`);
      form.setValue("leagueId", data.id);
      form.setValue("seasonId", data.season.ID);
    };

    if (leagueId) getLeague();
  }, [leagueId]);

  return (
    <Row>
      <Col xs={12} lg={8}>
        <Form.Label>League</Form.Label>
        <Controller
          name={"leagueId"}
          render={({ field }) => (
            <SelectLeague leagues={leagues} onChange={field.onChange} />
          )}
        />
      </Col>
      <Col xs={12} lg={8}>
        <Form.Label>Team</Form.Label>
        <Controller
          name={"teamId"}
          render={({ field }) => (
            <SelectTeam teams={teams} onChange={field.onChange} />
          )}
        />
      </Col>
      <Col xs={12} className="mt-3">
        <Button type="submit">Submit</Button>
      </Col>
    </Row>
  );
};

export default TeamForm;
