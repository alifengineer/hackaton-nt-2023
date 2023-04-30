import React from "react";
import { Form } from "react-bootstrap";
import { FormProvider, useForm } from "react-hook-form";
import { MatchForm } from "./components";
import axios from "../../utils/axios";
import toast from "react-hot-toast";

const CreateMatch = () => {
  const form = useForm();

  const onSubmit = async (formData) => {
    const payload = {
      home_team_id: formData.homeTeamId,
      away_team_id: formData.awayTeamId,
      league_id: formData.leagueId,
      m_date: formData.matchDate,
      tur: formData.tur,
    };
    const res = await axios.post("/admin-api/v1/match/", payload);
    if (res.status === 200 || res.status === 201) {
      toast.success("Match created");
    } else {
      toast.error("System Error");
    }
  };

  return (
    <FormProvider {...form}>
      <Form onSubmit={form.handleSubmit(onSubmit)}>
        <MatchForm />
      </Form>
    </FormProvider>
  );
};

export default CreateMatch;
