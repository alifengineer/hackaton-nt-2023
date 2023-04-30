import React from 'react';
import {Form} from "react-bootstrap";

const SelectTeam = ({teams, onChange}) => {
    return (
        <Form.Select onChange={(e) => console.log(e.target.value)}>
            {teams.map(option => (
                <option value={option.id} key={option.id}>{option.name}</option>
            ))}
        </Form.Select>
    );
};

export default SelectTeam;