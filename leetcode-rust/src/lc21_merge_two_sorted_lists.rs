// 21. Merge Two Sorted Lists
// Easy
// Topics
// Companies
// You are given the heads of two sorted linked lists list1 and list2.

// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

// Return the head of the merged linked list.

// Example 1:

// Input: list1 = [1,2,4], list2 = [1,3,4]
// Output: [1,1,2,3,4,4]
// Example 2:

// Input: list1 = [], list2 = []
// Output: []
// Example 3:

// Input: list1 = [], list2 = [0]
// Output: [0]

// Constraints:

// The number of nodes in both lists is in the range [0, 50].
// -100 <= Node.val <= 100
// Both list1 and list2 are sorted in non-decreasing order.

// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

use std::mem;
struct Solution;

impl Solution {
    pub fn merge_two_lists1(
        list1: Option<Box<ListNode>>,
        list2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        match (list1, list2) {
            (Some(l1), None) => Some(l1),
            (None, Some(l2)) => Some(l2),
            (None, None) => None,
            (Some(l1), Some(l2)) => match l1.val <= l2.val {
                true => Some(Box::new(ListNode {
                    val: l1.val,
                    next: Self::merge_two_lists1(l1.next, Some(l2)),
                })),
                false => Some(Box::new(ListNode {
                    val: l2.val,
                    next: Self::merge_two_lists1(Some(l1), l2.next),
                })),
            },
        }
    }

    // ### Recursive solution

    pub fn merge_two_lists2(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        match (l1, l2) {
            (None, None) => None,
            (Some(n), None) | (None, Some(n)) => Some(n),
            (Some(l1), Some(l2)) => {
                if l1.val >= l2.val {
                    Some(Box::new(ListNode {
                        val: l2.val,
                        next: Solution::merge_two_lists2(Some(l1), l2.next),
                    }))
                } else {
                    Some(Box::new(ListNode {
                        val: l1.val,
                        next: Solution::merge_two_lists2(l1.next, Some(l2)),
                    }))
                }
            }
        }
    }

    // ### Other Solution

    #[allow(dead_code)]
    pub fn merge_two_lists3(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let mut result = l1;
        let mut l2 = l2;
        let mut lsmall = &mut result;
        let lbig = &mut l2;
        while lbig.is_some() {
            if lsmall.is_none() || lsmall.as_ref()?.val > lbig.as_ref()?.val {
                mem::swap(lsmall, lbig);
            }
            if lsmall.is_some() {
                lsmall = &mut lsmall.as_mut()?.next;
            }
        }
        result
    }

    // ### other solution

    pub fn merge_two_lists4(
        mut a_ptr: Option<Box<ListNode>>,
        mut b_ptr: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        // use std::mem;

        // We create two vars which are pointing to the same place in memory. While `ptr` will be continuously updated,
        // `head` will not and will keep pointing to the location of the first node.
        let mut head = None;
        let mut ptr = &mut head;

        while a_ptr.is_some() && b_ptr.is_some() {
            let ap = &mut a_ptr;
            let bp = &mut b_ptr;

            // We set `tmp` to point to the next target node we want to append.
            let mut tmp = if bp.as_ref().unwrap().val > ap.as_ref().unwrap().val {
                ap
            } else {
                bp
            };

            // We swap the data in `ptr` (the currently built up list) and the data in `tmp` (the next target node), moving the
            // currently built up list to `tmp`. `ptr` now points to the next target node.
            mem::swap(ptr, &mut tmp);

            // We swap the data in `tmp` (the currently built up list) with the data of `ptr.next`, effectively clipping
            // off the target node's tail. `tmp` now holds the tail, `ptr.next` now (only) holds the next target node.
            mem::swap(tmp, &mut ptr.as_mut().unwrap().next);

            // All the while, we only swapped what was held in memory, not what the variables were pointing to. We now move
            // the cursor along to its current `next`, the target node that we just appended.
            ptr = &mut ptr.as_mut().unwrap().next;
        }

        mem::swap(
            ptr,
            if a_ptr.is_none() {
                &mut b_ptr
            } else {
                &mut a_ptr
            },
        );

        head
    }
}

pub fn main() {
    println! {"Merge two sorted lists..."};
}
