-- +goose Up
CREATE TABLE jobs (
    postionId SERIAL PRIMARY KEY,
    positionName VARCHAR(255) NOT NULL,
    companyId INT NOT NULL REFERENCES companies(companyId),
    experienceRequired VARCHAR(255) NOT NULL,
    expectedCompensation VARCHAR(255) NOT NULL,
    employmentType VARCHAR(255) NOT NULL,
    jobLocationType VARCHAR(255) NOT NULL
);



-- +goose Down
DROP TABLE jobs;