import React from 'react';
import ReactSelect from "react-select";

const SelectLeague = ({leagues, onChange}) => {
    return (
        <ReactSelect options={leagues}
                     onChange={(option) => onChange(option ? option.id : null)}
                     getOptionLabel={option => option.name}
                     getOptionValue={option => option.id}
                     isClearable
        />
    );
};

export default SelectLeague;