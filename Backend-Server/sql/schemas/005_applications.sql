-- +goose Up
CREATE TABLE applications (
    id SERIAL PRIMARY KEY,
    department VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    companyid INT NOT NULL REFERENCES companies(companyId),
    positionid INT NOT NULL REFERENCES jobs(postionId),
    interviewId INT REFERENCES Interviews(id),
    userid INT NOT NULL REFERENCES accounts(id)
);



-- +goose Down
DROP TABLE applications;
