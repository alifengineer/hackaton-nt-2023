import React, {useEffect, useState} from 'react';
import {Col, Form, Row} from "react-bootstrap";
import SelectTeam from "./SelectTeam";
import SelectLeague from "./SelectLeague";
import axios from "../../../utils/axios";

const MatchForm = () => {
    const [teams, setTeams] = useState([]);
    const [leagues, setLeagues] = useState([]);

    useEffect(() => {
        const getAllTeams = async () => {
            const {data: {data: {teams}}} = await axios.get("/api/v1/team/");
            setTeams(teams);
        }

        const getAllLeagues = async () => {
            const {data: {data}} = await axios.get("api/v1/league/");
            setLeagues(data);
        }

        getAllTeams();
        getAllLeagues();
    }, [])

    console.log(teams, leagues);

    return (
        <Row>
            <Col xs={12}>
                <Form.Group>
                    <Form.Label>Home team</Form.Label>
                    <SelectTeam teams={teams}/>
                </Form.Group>
            </Col>
            <Col xs={12}>
                <Form.Group>
                    <Form.Label>Away team</Form.Label>
                    <SelectTeam teams={teams}/>
                </Form.Group>
            </Col>
            <Col xs={12}>
                <Form.Group>
                    <Form.Label>League</Form.Label>
                    <SelectLeague leagues={leagues}/>
                </Form.Group>
            </Col>
            <Col xs={12}>
                <Form.Group>
                    <Form.Label>Match date</Form.Label>
                    <Form.Control type="date"
                                  onChange={(e) => console.log(e.target.valueAsDate)}
                    />
                </Form.Group>
            </Col>
            <Col xs={12}>
                <Form.Group>
                    <Form.Label>Tur</Form.Label>
                </Form.Group>
            </Col>
        </Row>
    );
};

export default MatchForm;