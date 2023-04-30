import React from 'react';
import {Form} from "react-bootstrap";

const SelectLeague = ({leagues, onChange}) => {
    return (
        <Form.Select onChange={(e) => console.log(e.target.value)}>
            {leagues.map(option => (
                <option value={option.id} key={option.id}>{option.name}</option>
            ))}
        </Form.Select>
    );
};

export default SelectLeague;