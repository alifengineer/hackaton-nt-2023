import React from 'react';
import {Form} from "react-bootstrap";
import {FormProvider, useForm} from "react-hook-form";
import {MatchForm} from "./components";

const CreateMatch = () => {
    const form = useForm();

    const onSubmit = (formData) => {

    }

    return (
        <FormProvider {...form}>
            <Form onSubmit={form.handleSubmit(onSubmit)}>
                <MatchForm />
            </Form>
        </FormProvider>
    );
};

export default CreateMatch;