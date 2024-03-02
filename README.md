# A specification document for the HRMS

### A collection of interfaces to where our system will be depending on

## Table of contents

1. Terminologies
2. A General Rule of Thumb

## Terminologies

Human Resource Management System (HRMS), is a software designed to automate various HR functions within an organization. **It reduces the manual labor in managing employee data, payroll, benefits administration, recruitment, performance management, time tracking, and other functions related to human resource.**

## A General Rule of Thumb

Each interfaces that will be described in this source must be followed accordingly so that there will be no confusions throughout the development cycle and helps make your HRMS robust.

Below are the types, enums and interfaces that will be outlined:

1. Worker
2. WorkerIdentityCard
3. WorkerIdentityCardRelation
4. Organization
5. Member
6. WorkerOganizationRelation
7. JoinRequest
8. OrganizationPendingRequestRelation
9. WorkerPendingRequestRelation
10. Role
11. RoleShiftRelation
12. WorkerRoleRelation
13. Team
14. TeamMemberRelation
15. Compensation
16. Addition
17. Deduction
18. CompensationAdditionRelation
19. CompensationDeductionRelation
20. Payroll
21. PayrollCompensationRelation
22. Attendance
23. Shift
24. OverrideShift
25. VerifiedWorkerIdentityCard