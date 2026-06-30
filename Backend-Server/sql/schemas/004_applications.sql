-- +goose Up
CREATE TABLE applications (
    id SERIAL PRIMARY KEY,
    department VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    company INT NOT NULL REFERENCES companies(companyId),
    position INT NOT NULL REFERENCES jobs(postionId),
    interviewId INT REFERENCES Interviews(id)
);



-- +goose Down
DROP TABLE applications;
