import React, {useEffect, useState} from 'react';
import {Button, Col, Form, Row} from "react-bootstrap";
import SelectTeam from "./SelectTeam";
import SelectLeague from "./SelectLeague";
import axios from "../../../utils/axios";
import {Controller, useFormContext, useWatch} from "react-hook-form";
import DatePicker from "react-datepicker";

const MatchForm = () => {
    const [teams, setTeams] = useState([]);
    const [leagues, setLeagues] = useState([]);
    const [league, setLeague] = useState(null);

    const [turs, setTurs] = useState([]);
    const form = useFormContext();

    const homeTeamId = useWatch({name: 'homeTeamId'})
    const awayTeamId = useWatch({name: 'awayTeamId'});
    const leagueId = useWatch({name: 'leagueId'});


    useEffect(() => {
        const getAllLeagues = async () => {
            const {data: {data}} = await axios.get("api/v1/league/");
            setLeagues(data);
        }

        getAllLeagues();
    }, [])

    useEffect(() => {
        const getLeague = async () => {
            const {data: {data}} = await axios.get(`api/v1/league/${leagueId}`);
            setLeague(data);

            const options = [];
            for (let i = 1; i < data.tur.soni; i++) {
                options.push(i);
            }
            setTurs(options);
        }

        if (leagueId)
            getLeague();
    }, [leagueId])

    useEffect(() => {
        const getAllTeams = async () => {
            const {data: {data: {teams}}} = await axios.get(`/api/v1/league/${league.id}/seasons/${league.season.ID}/teams`);
            setTeams(teams);
        }

        if (league)
            getAllTeams();
    }, [league])

    return (
        <Row>
            <Col xs={12} lg={3}>
                <Form.Group>
                    <Form.Label>League</Form.Label>
                    <Controller name={'leagueId'}
                                render={({field}) => (
                                    <SelectLeague leagues={leagues}
                                                  onChange={field.onChange}
                                    />
                                )}
                    />
                </Form.Group>
            </Col>
            <Col xs={12} lg={3}>
                <Form.Group>
                    <Form.Label>Home team</Form.Label>
                    <Controller name={'homeTeamId'}
                                render={({field}) => {
                                    const options = teams.filter(x => x.id !== awayTeamId);
                                    return (
                                        <SelectTeam teams={options}
                                                    onChange={field.onChange}
                                        />
                                    )
                                }}
                    />

                </Form.Group>
            </Col>
            <Col xs={12} lg={3}>
                <Form.Group>
                    <Form.Label>Away team</Form.Label>
                    <Controller name={'awayTeamId'}
                                render={({field}) => {
                                    const options = teams.filter(x => x.id !== homeTeamId);
                                    return (
                                        <SelectTeam teams={options}
                                                    onChange={field.onChange}
                                        />
                                    )
                                }}
                    />
                </Form.Group>
            </Col>
            <Col xs={12} lg={3}>
                <Form.Group>
                    <Form.Label>Tur</Form.Label>
                    <Form.Select {...form.register('tur')}>
                        {turs.map((tur, index) => (
                            <option value={tur} key={tur}>{tur}</option>
                        ))}
                    </Form.Select>
                </Form.Group>
            </Col>
            <Col xs={12} lg={3}>
                <Form.Group>
                    <Form.Label>Match date</Form.Label>
                    <Controller
                        name="matchDate"
                        render={({field}) => (
                            <DatePicker
                                isClearable
                                className='form-control'
                                selected={field.value}
                                onChange={date => field.onChange(date)}
                            />
                        )}
                    />
                </Form.Group>
            </Col>
            <Col xs={12} lg={3} className="mt-3 mt-lg-0">
                <Button type="submit" className="mt-lg-4">Submit</Button>
            </Col>
        </Row>
    );
};

export default MatchForm;