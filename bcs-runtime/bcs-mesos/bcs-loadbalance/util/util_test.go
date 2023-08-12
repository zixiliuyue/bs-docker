

package util_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-loadbalance/util"
)

var _ = Describe("Util", func() {
	Describe("Execute command", func() {
		It("Normal test", func() {
			result, flag := ExeCommand("echo hello")
			Expect(result).To(Equal("hello\n"))
			Expect(flag).To(Equal(true))
		})
		It("Execute error test", func() {
			_, flag := ExeCommand("./unknown-command")
			Expect(flag).To(Equal(false))
		})
	})

	Describe("GetSubsection", func() {
		It("[a,b,c] - [b,c,d] -> [a]", func() {
			subs := GetSubsection([]string{"a", "b", "c"}, []string{"b", "c", "d"})
			Expect(subs).To(Equal([]string{"a"}))
		})
		It("[a,b,c] - [] -> [a,b,c]", func() {
			subs := GetSubsection([]string{"a", "b", "c"}, []string{})
			Expect(subs).To(Equal([]string{"a", "b", "c"}))
		})
		It("[] - [a,b,c] -> [] ", func() {
			subs := GetSubsection([]string{}, []string{"a", "b", "c"})
			Expect(subs).To(Equal([]string{}))
		})
	})

	Describe("TrimSpecialChar", func() {
		It("a/b c~d*e.f\\g -> abcdefd", func() {
			Expect(TrimSpecialChar("a/b c~d*e.f\\g")).To(Equal("abcdefg"))
		})
	})

	Describe("GetValidZookeeperPath", func() {
		It("a/b c~d*e.f\\g -> a_bcdef", func() {
			Expect(GetValidZookeeperPath("a/b c~d*e.f\\g")).To(Equal("a_bcdefg"))
		})
		It("/ -> ", func() {
			Expect(GetValidZookeeperPath("/")).To(Equal(""))
		})
	})

	Describe("GetValidTargetGroupSub", func() {
		It("a/b c~d*e.f\\g -> a-bcdef", func() {
			Expect(GetValidTargetGroupSub("a/b c~d*e.f\\g")).To(Equal("a-bcdefg"))
		})
		It("/ -> ", func() {
			Expect(GetValidTargetGroupSub("/")).To(Equal(""))
		})
	})
})
