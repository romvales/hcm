# A specification document for the HRMS

### A collection of interfaces to where our system will be depending on

## Table of contents

1. Terminologies
2. A General Rule of Thumb

## Terminologies

Human Resource Management System (HRMS), is a software designed to automate various HR functions within an organization. **It reduces the manual labor in managing employee data, payroll, benefits administration, recruitment, performance management, time tracking, and other functions related to human resource.**

## A General Rule of Thumb

Each interfaces that will be described in this source must be followed accordingly so that there will be no confusions throughout the development cycle and helps make your HRMS robust.

Below are the types, enums and interfaces that will be described:

1. Worker

- Role
- Team
- Worker
- WorkerIdentityCard
- WorkerAddress
- enum:WorkerAddressType
- enum:WorkerGender
- enum:WorkerStatus
- enum:WorkerType
- enum:TeamStatus
- enum:RoleStatus

2. Shifts

- StandardShift
- OverrideShift
- enum:StandardShiftDay

3. Attendance

- Attendance
- enum:AttendanceType
- enum:AttendanceStatus

4. Payroll


5. Compute


6. Compensation

- WorkerPayInfo
- WorkerPayInfoOverride
- Compensation
- Addition
- Deduction
- enum:AdditionStatus
- enum:AdditionType
- enum:DeductionStatus
- enum:DeductionType
- enum:WorkerPayInfoType
