import React from 'react';

type Props = {
    imagePortraitUrl: string;
    fullName: string;
    officeLocation: string;
    linkedInHandle: string;
    githubHandle: string;
    twitterHandle: string;
};

export const EmployeeCard: React.FunctionComponent<Props> = (props: Props) => {
    return (
        <div className="employee-card">
            <div className="photo"><img src={props.imagePortraitUrl} alt="Portrait" /></div>
            <div className="info">
                <div className="name-location">
                    <div>{props.fullName}</div>
                    <div>Office: {props.officeLocation}</div>
                </div>
                <div className="socials">
                    {!!props.linkedInHandle && <a href={`https://linkedin.com/in/${props.linkedInHandle}`}><img src="/images/linkedin.png" alt="LinkedIn" /></a>}
                    {!!props.githubHandle && <a href={`https://github.com/${props.githubHandle}`}><img src="/images/github.png" alt="GitHub" /></a>}
                    {!!props.twitterHandle && <a href={`https://twitter.com/${props.twitterHandle}`}><img src="/images/twitter.png" alt="Twitter" /></a>}
                </div>
            </div>
        </div>
    );
};