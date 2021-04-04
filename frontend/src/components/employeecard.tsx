import React from 'react';

type Props = {
    fullName: string;
    officeLocation: string;
    linkedInHandle: string;
    githubHandle: string;
    twitterHandle: string;
};

export const EmployeeCard: React.FunctionComponent<Props> = (props: Props) => {
    return (
        <div>
            <div>{props.fullName}</div>
            <div>Office: {props.fullName}</div>
            {!!props.linkedInHandle && <div>LI: {props.linkedInHandle}</div>}
            {!!props.githubHandle && <div>GH: {props.githubHandle}</div>}
            {!!props.twitterHandle && <div>TW: {props.twitterHandle}</div>}
            <br />
        </div>
    );
};