import unittest
import day1

class TestDay1(unittest.TestCase):
    def test1(self):
        got = day1.calculate_fuel(12)
        self.assertEqual(got, 2)
    def test2(self):
        got = day1.calculate_fuel(14)
        self.assertEqual(got, 2)
    def test3(self):
        got = day1.calculate_fuel(1969)
        self.assertEqual(got, 654)
    def test4(self):
        got = day1.calculate_fuel(100756)
        self.assertEqual(got, 33583)


