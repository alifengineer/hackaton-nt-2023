import React from 'react';
import ReactSelect from "react-select";

const SelectTeam = ({teams, onChange}) => {
    return (
        <ReactSelect options={teams}
                     onChange={(option) => onChange(option ? option.id : null)}
                     getOptionLabel={option => option.name}
                     getOptionValue={option => option.id}
                     isClearable
        />
    );
};

export default SelectTeam;