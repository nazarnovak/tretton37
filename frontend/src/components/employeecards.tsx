import React, { useEffect, useState } from 'react';
import axios from 'axios';

import { EmployeeCard } from 'components/employeecard';

import { Employee } from 'types';

const apiGetAllEmployees = '/api/employees';

export const EmployeeCards: React.FunctionComponent = () => {
    const [isLoading, setIsLoading] = useState(false);
    const [isError, setIsError] = useState(false);
    const [employees, setEmployees] = useState([]);

    useEffect(()  => {
        // Not sure if possible to extract this into a separate function outside of useEffect
        const fetchData = async () => {
            setIsError(false);
            setIsLoading(true);
        
            await axios.get(apiGetAllEmployees)
            .then(res => {
                setEmployees(res.data);
            })
            .catch(err => {
                console.log(err);
                setIsError(true);
            });

            setIsLoading(false);
        };

        fetchData();
    }, []);

    const getEmployeeCards = (employees: Employee[]) => {
        if (!employees.length) {
            return (<div>Currently there are no employees</div>);
        }

        return employees.map((employee: Employee, index: number) => {
            return <EmployeeCard 
                key={index} 
                fullName={employee.name} 
                officeLocation={employee.office}
                linkedInHandle={employee.linkedin} 
                githubHandle={employee.github}
                twitterHandle={employee.twitter}
            />;
        });
    };

    return (
        <div>
            {isLoading && <h3>Loading data...</h3>}
            {!isLoading && isError && <h3>Something went wrong, please check console for more details</h3>}
            {!isLoading && !isError && getEmployeeCards(employees)}
        </div>
    );    
};
