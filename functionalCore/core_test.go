package functionalcore_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	fc "nilsherzig.com/imperativeShell/functionalCore"
)

var _ = Describe("Core", func() {
	var (
		members     fc.Members
		memberSlice []fc.Member
	)

	BeforeEach(func() {
		memberSlice = []fc.Member{
			{Name: "John", IsReady: true},
			{Name: "Jane", IsReady: false},
		}

		members = fc.New(memberSlice)
	})

	It("should return a new Members object", func() {
		Expect(len(members.Items)).To(Equal(2))
		Expect(members.Items).To(Equal(memberSlice))
	})

	It("should filter ready members", func() {
		Expect(len(fc.ReadyMembers(members).Items)).
			To(Equal(1))
		Expect(fc.ReadyMembers(members).Items).
			To(Equal([]fc.Member{{Name: "John", IsReady: true}}))
	})

	It("should add a new member", func() {
		newMember := fc.Member{Name: "Jack", IsReady: true}

		Expect(fc.AddMember(members, newMember.Name, newMember.IsReady).Items).
			To(Equal(append(memberSlice, newMember)))
	})
})
