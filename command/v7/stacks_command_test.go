package v7_test

import (
	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/command/commandfakes"
	"code.cloudfoundry.org/cli/command/translatableerror"
	. "code.cloudfoundry.org/cli/command/v7"
	"code.cloudfoundry.org/cli/command/v7/v7fakes"
	"code.cloudfoundry.org/cli/util/ui"
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
)

var _ = Describe("stacks Command", func() {
	var (
		cmd             StacksCommand
		testUI          *ui.UI
		fakeConfig      *commandfakes.FakeConfig
		fakeSharedActor *commandfakes.FakeSharedActor
		fakeActor       *v7fakes.FakeStacksActor
		executeErr      error
		args            []string
	)

	const tableHeaders = `name\s+description`

	JustBeforeEach(func() {
		executeErr = cmd.Execute(args)
	})

	BeforeEach(func() {
		testUI = ui.NewTestUI(nil, NewBuffer(), NewBuffer())
		fakeConfig = new(commandfakes.FakeConfig)
		fakeSharedActor = new(commandfakes.FakeSharedActor)
		fakeActor = new(v7fakes.FakeStacksActor)
		args = nil

		cmd = StacksCommand{
			UI:          testUI,
			Config:      fakeConfig,
			SharedActor: fakeSharedActor,
			Actor:       fakeActor,
		}
	})

	When("too many args are passed", func() {
		BeforeEach(func() {
			args = []string{"first-extra-arg", "second-extra-arg"}
		})

		It("returns a TooManyArgumentsError", func() {
			Expect(executeErr).To(MatchError(
				translatableerror.TooManyArgumentsError{
					ExtraArgument: "first-extra-arg",
				},
			))
		})
	})

	When("StacksActor returns an error", func() {
		var expectedErr error
		BeforeEach(func() {
			warnings := v7action.Warnings{"warning-1", "warning-2"}
			expectedErr = errors.New("some-error")
			fakeActor.GetStacksReturns(nil, warnings, expectedErr)
		})

		It("prints that error with warnings", func() {
			Expect(executeErr).To(Equal(expectedErr))
			Expect(testUI.Err).To(Say("warning-1"))
			Expect(testUI.Err).To(Say("warning-2"))
			Expect(testUI.Out).ToNot(Say(tableHeaders))
		})
	})

	When("everything is perfect", func() {
		BeforeEach(func() {
			stacks := []v7action.Stack{
				{Name: "stack1", Description: "desc1"},
				{Name: "stack2", Description: "desc2"},
			}
			fakeActor.GetStacksReturns(stacks, v7action.Warnings{"warning-1", "warning-2"}, nil)
		})

		It("asks the StacksActor for a list of stacks", func() {
			Expect(fakeActor.GetStacksCallCount()).To(Equal(1))
		})

		It("prints warnings", func() {
			Expect(testUI.Err).To(Say(`warning-1`))
			Expect(testUI.Err).To(Say(`warning-2`))
		})

		It("prints the list of stacks", func() {
			Expect(testUI.Out).To(Say(tableHeaders))
			Expect(testUI.Out).To(Say(`stack1\s+desc1`))
			Expect(testUI.Out).To(Say(`stack2\s+desc2`))
		})
	})
})
