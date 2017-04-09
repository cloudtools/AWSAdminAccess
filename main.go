package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

var policyArn = "arn:aws:iam::aws:policy/AdministratorAccess"
var assumeRoleDocument = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::%s:root"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}`

func main() {
	var account, role string

	flag.StringVar(&account, "a", "", "Account to trust")
	flag.StringVar(&role, "r", "", "Role name to create")
	flag.Parse()

	if account == "" {
		fmt.Printf("Must specify account option\n")
		os.Exit(0)
	}

	if role == "" {
		fmt.Printf("Must specify role name option\n")
		os.Exit(0)
	}

	sess := session.Must(session.NewSession())
	svc := iam.New(sess)

	// First see if the role already exists
	params := &iam.GetRoleInput{
		RoleName: aws.String(role),
	}
	_, err := svc.GetRole(params)
	if err == nil {
		fmt.Printf("Role %s already exists\n", role)
		os.Exit(0)
	}

	// Make sure the policy exists before creating the role
	policyParams := &iam.GetPolicyInput{
		PolicyArn: aws.String(policyArn),
	}
	_, err = svc.GetPolicy(policyParams)
	if err != nil {
		fmt.Printf("Policy %s does not exist: %v\n", policyArn, err)
		os.Exit(0)
	}

	// Create the role with a trust policy document
	assumeRoleString := fmt.Sprintf(assumeRoleDocument, account)
	roleParams := &iam.CreateRoleInput{
		AssumeRolePolicyDocument: aws.String(assumeRoleString),
		RoleName:                 aws.String(role),
	}
	roleOutput, err := svc.CreateRole(roleParams)
	if err != nil {
		fmt.Printf("Cannot create role %s\n", err)
		os.Exit(0)
	}

	// Attach the role policy onto the role
	_, err = svc.AttachRolePolicy(&iam.AttachRolePolicyInput{
		PolicyArn: aws.String(policyArn),
		RoleName:  aws.String(role),
	})
	if err != nil {
		fmt.Printf("AttachRolePolicy failed: %v\n", err)
		os.Exit(0)
	}

	fmt.Printf("Role %s created - ARN: %s\n", role, *roleOutput.Role.Arn)
}
