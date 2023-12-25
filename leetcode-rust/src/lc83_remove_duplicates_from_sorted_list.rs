/*
83. Remove Duplicates from Sorted List
Easy
Topics
Companies
Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

Example 1:

Input: head = [1,1,2]
Output: [1,2]
Example 2:

Input: head = [1,1,2,3,3]
Output: [1,2,3]

Constraints:

The number of nodes in the list is in the range [0, 300].
-100 <= Node.val <= 100
The list is guaranteed to be sorted in ascending order.
*/
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

pub fn delete_duplicates1(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    if head.is_none() {
        return None;
    }
    let mut h = head;
    let mut p1 = h.as_mut().unwrap();
    while let Some(p2) = p1.next.as_mut() {
        if p1.val == p2.val {
            p1.next = p2.next.take();
        } else {
            p1 = p1.next.as_mut().unwrap();
        }
    }
    h
}

// ### other soln

pub fn delete_duplicates2(mut head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    let mut curr_opt = head.as_mut();

    while let Some(curr) = curr_opt {
        let mut next_opt = curr.next.take();

        while let Some(next) = next_opt.as_mut() {
            if next.val == curr.val {
                next_opt = next.next.take();
            } else {
                curr.next = next_opt;
                break;
            }
        }
        curr_opt = curr.next.as_mut();
    }
    head
}

pub fn main() {
    println! {"Remove duplicates from sorted list"};
}
