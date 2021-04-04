import React from 'react';

import { EmployeeCards } from 'components/employeecards';

import './App.css';

export const App: React.FunctionComponent = () => {
  return (
    <div className="App">
      <h2>The fellowship of the tretton37</h2>
      <EmployeeCards />
    </div>
  );
};
