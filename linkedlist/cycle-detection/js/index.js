class ListNode {
  constructor(val, next = null) {
    this.val = val;
    this.next = next;
  }
}

function findCyclePoint(head) {

  let hale = head;
  let tortoise = head;

  while(true) {
    if (hale.next === null || tortoise.next === null || tortoise.next.next === null) {
      return null;
    }
    
    hale = hale.next;
    tortoise = tortoise.next.next;

    if (hale === tortoise) {
      break;
    }

  }

  let p1 = head;
  let p2 = tortoise;

  while(true) {
    p1 = p1.next;
    p2 = p2.next;

    if (p1 === p2) {
      return p1;
    }

  }

  return null;

}

module.exports = {
  ListNode,
  findCyclePoint,
};