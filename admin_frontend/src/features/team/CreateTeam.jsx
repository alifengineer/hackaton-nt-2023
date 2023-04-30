import React from "react";
import { FormProvider, useForm } from "react-hook-form";
import axios from "../../utils/axios";
import { Form } from "react-bootstrap";
import { TeamForm } from "./components";

const CreateTeam = () => {
  const form = useForm();

  const onSubmit = async (formData) => {
    const teams = [];
    teams.push({ id: formData.teamId });
    const payload = {
      teams: teams,
      league_id: formData.leagueId,
      season_id: formData.seasonId,
    };

    await axios.post(
      `/admin-api/v1/league/${formData.leagueId}/seasons/${formData.seasonId}/teams`,
      payload
    );
  };

  return (
    <FormProvider {...form}>
      <Form onSubmit={form.handleSubmit(onSubmit)}>
        <TeamForm />
      </Form>
    </FormProvider>
  );
};

export default CreateTeam;
