# Copyright 2016 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import textwrap

from common.dependency import Dependency
from crash.callstack import StackFrame, CallStack
from crash.fracas_parser import FracasParser
from crash.stacktrace import Stacktrace
from crash.test.stacktrace_test_suite import StacktraceTestSuite
from crash.type_enums import CallStackFormatType


class FracasParserTest(StacktraceTestSuite):

  def testFracasParserIsStartOfNewCallSTack(self):
    parser = FracasParser()
    self.assertEqual(parser._IsStartOfNewCallStack('dummy line'),
                     (False, None, None))
    self.assertEqual(parser._IsStartOfNewCallStack('CRASHED [EXC @ 0x508]'),
                     (True, 0, CallStackFormatType.DEFAULT))

  def testReturnEmptyStacktraceForEmptyString(self):
    parser = FracasParser()
    deps = {'src/': Dependency('src/', 'https://repo', '1')}

    self._VerifyTwoStacktracesEqual(parser.Parse('', deps),
                                    Stacktrace())

  def testFracasParserParseLineMalformatedCallstack(self):
    parser = FracasParser()
    deps = {'src/': Dependency('src/', 'https://repo', '1')}
    stacktrace_string = textwrap.dedent(
        """
        CRASHED [EXC @ 0x508]
        #0 [RESTRICTED]
        #1 [RESTRICTED]
        """
    )
    self._VerifyTwoStacktracesEqual(parser.Parse(stacktrace_string, deps),
                                    Stacktrace())

  def testFracasParserParseLineOneCallstack(self):
    parser = FracasParser()
    deps = {'src/': Dependency('src/', 'https://repo', '1')}
    stacktrace_string = textwrap.dedent(
        """
        CRASHED [EXC @ 0x508]
        #0 0x7fee in a::c(p* &d) src/f0.cc:177
        #1 0x4b6e in a::d(a* c) src/f1.cc:227
        #2 0x7ff9 in a::e(int) src/f2.cc:87:1
        """
    )

    stacktrace = parser.Parse(stacktrace_string, deps)

    expected_callstack = CallStack(0)
    expected_callstack.extend(
        [StackFrame(0, 'src/', 'a::c(p* &d)', 'f0.cc', 'src/f0.cc', [177]),
         StackFrame(1, 'src/', 'a::d(a* c)', 'f1.cc', 'src/f1.cc', [227]),
         StackFrame(2, 'src/', 'a::e(int)', 'f2.cc', 'src/f2.cc', [87, 88])])

    expected_stacktrace = Stacktrace()
    expected_stacktrace.append(expected_callstack)

    self._VerifyTwoStacktracesEqual(stacktrace, expected_stacktrace)

  def testFracasParserParseLineMultipleCallstacks(self):
    parser = FracasParser()
    deps = {'src/': Dependency('src/', 'https://repo', '1')}
    stacktrace_string = textwrap.dedent(
        """
        CRASHED [EXC @ 0x66]
        #0 0x7fee in a::b::c(p* &d) src/f0.cc:177
        #1 0x4b6e in a::b::d(a* c) src/f1.cc:227

        CRASHED [EXC @ 0x508]
        #0 0x8fee in e::f::g(p* &d) src/f.cc:20:2
        #1 0x1fae in h::i::j(p* &d) src/ff.cc:9:1
        """
    )

    stacktrace = parser.Parse(stacktrace_string, deps)

    expected_callstack0 = CallStack(0)
    expected_callstack0.extend(
        [StackFrame(0, 'src/', 'a::b::c(p* &d)', 'f0.cc', 'src/f0.cc', [177]),
         StackFrame(1, 'src/', 'a::b::d(a* c)', 'f1.cc', 'src/f1.cc', [227])])

    expected_callstack1 = CallStack(0)
    expected_callstack1.extend(
        [StackFrame(
            0, 'src/', 'e::f::g(p* &d)', 'f.cc', 'src/f.cc', [20, 21, 22]),
         StackFrame(
             1, 'src/', 'h::i::j(p* &d)', 'ff.cc', 'src/ff.cc', [9, 10])])

    expected_stacktrace = Stacktrace()
    expected_stacktrace.extend([expected_callstack0, expected_callstack1])

    self._VerifyTwoStacktracesEqual(stacktrace, expected_stacktrace)
