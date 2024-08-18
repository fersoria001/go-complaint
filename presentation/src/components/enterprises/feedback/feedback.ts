import { User } from "@/gql/graphql";
import { ComplaintReply, ReplyReview } from "@/gql/graphql";

export type ColorSquare = {
  selected: boolean;
  picked: boolean;
  fixed: boolean;
  color: string;
};

export const defaultSquares = [
  {
    selected: false,
    picked: false,
    fixed: false,
    color: "#ef4444",
  },
  {
    selected: false,
    picked: false,
    fixed: false,
    color: "#eab308",
  },
  {
    selected: false,
    picked: false,
    fixed: false,
    color: "#3b82f6",
  },
];

export function getSelectedColor(colorSquares: ColorSquare[]) {
  return colorSquares.find((v) => v.selected === true);
}

/**
 *
 * @param selectedColor
 * @param colorSquares
 * @returns {ColorSquare[]}
 * @throws {Error} if color not found in colorSquares
 * @throws {ColorAlreadyPickedError} if color already picked
 */
export const selectAColor = (
  selectedColor: string,
  colorSquares: ColorSquare[]
): ColorSquare[] => {
  const square = colorSquares.find((square) => square.color === selectedColor);
  if (!square) {
    throw new Error(`Color ${selectedColor} not found in colorSquares`);
  }
  if (square.fixed) {
    throw new ColorsFixedError(
      `finish your review by adding a comment first or delete it`
    );
  }
  if (square.picked) {
    console.log("squares inside selectAColor", colorSquares);
    throw new ColorAlreadyPickedError(`${selectedColor} color already picked`);
  }
  const newSquare = { ...square, selected: true };
  const newSquares = colorSquares.map((v) =>
    v.color === selectedColor ? newSquare : { ...v, selected: false }
  );
  return newSquares;
};

/**
 *
 * @param colorSquares
 * @param color
 * @returns
 * @throws {Error} if color not found in colorSquares
 */
export const unpickColor = (
  colorSquares: ColorSquare[],
  color: string
): ColorSquare[] => {
  const squareIndex = colorSquares.findIndex(
    (square) => square.color === color
  );
  if (squareIndex !== -1) {
    const newSquares = colorSquares.map((square, idx) =>
      idx === squareIndex
        ? {
            ...square,
            selected: false,
            picked: false,
          }
        : square
    );
    return newSquares;
  }
  throw new Error(`Color ${color} not found in colorSquares`);
};

export const colorAlreadyPicked = (
  colorSquares: ColorSquare[],
  color: string
): boolean => {
  const squareIndex = colorSquares.findIndex(
    (square) => square.color === color
  );
  if (squareIndex !== -1) {
    const square = colorSquares[squareIndex];
    if (square.picked) {
      return true;
    }
  }
  return false;
};

/**
 *
 * @param colorSquares
 * @param color
 * @returns
 * @throws {Error} if color not found in colorSquares
 * @throws {ColorAlreadyPickedError} if color already picked
 */
export const colorPicked = (
  colorSquares: ColorSquare[],
  color: string
): ColorSquare[] => {
  const squareIndex = colorSquares.findIndex(
    (square) => square.color === color
  );
  if (squareIndex !== -1) {
    const square = colorSquares[squareIndex];
    if (square.picked) {
      throw new ColorAlreadyPickedError(`${color} color already picked`);
    }
    const newSquares = colorSquares.map((square, idx) =>
      idx === squareIndex
        ? {
            ...square,
            selected: false,
            picked: true,
          }
        : square
    );
    return newSquares;
  }
  throw new Error(`Color ${color} not found in colorSquares`);
};

export const fixSquaresPos = (
  colorSquares: ColorSquare[],
  value = true
): ColorSquare[] => {
  const newSquares = colorSquares.map((square) => {
    return { ...square, fixed: value };
  });
  return newSquares;
};

export type FeedbackErrorType = { [key: string]: { [key: string]: string } };

export type ColorError = {
  colorKey: string;
  error: string;
};

export class ColorAlreadyPickedError extends Error {
  constructor(message: string) {
    super(message);
    this.name = "ColorAlreadyPickedError";
  }
}

export class ColorsFixedError extends Error {
  constructor(message: string) {
    super(message);
    this.name = "ColorsFixedError";
  }
}

export type CommentObject = {
  comment: string;
  color: string;
  done: boolean;
  repliesRefs: Map<string, React.RefObject<HTMLLIElement>>;
  selectedReplies: ComplaintReply[];
  reviewer: Partial<User>;
};

export const createComment = (
  color: string,
  reviewer: Partial<User>
): CommentObject => {
  return {
    reviewer: reviewer,
    color: color,
    comment: "",
    done: false,
    repliesRefs: new Map(),
    selectedReplies: [],
  };
};

export const setReviewer = (
  comment: CommentObject,
  reviewer: Partial<User>
): CommentObject => {
  return { ...comment, reviewer: reviewer };
};

export const setBody = (
  comment: CommentObject,
  body: string
): CommentObject => {
  return { ...comment, comment: body };
};

export const setDone = (
  comment: CommentObject,
  done: boolean
): CommentObject => {
  return { ...comment, done: done };
};

export type IDOMRect = {
  left: number;
  right: number;
  top: number;
  bottom: number;
  width: number;
  height: number;
};

export const currentDOMRect = (ref: React.RefObject<any> | null): IDOMRect => {
  if (ref && ref.current) {
    const rect = ref.current.getBoundingClientRect();
    return {
      left: rect.left,
      right: rect.right,
      top: rect.top,
      bottom: rect.bottom,
      width: rect.width,
      height: rect.height,
    };
  }
  return {
    left: 0,
    right: 0,
    top: 0,
    bottom: 0,
    width: 0,
    height: 0,
  };
};

const offsetScrollY = (topValue: number) => {
  const scrollTop = window.scrollY || document.documentElement.scrollTop;
  return topValue + scrollTop;
};

const limitScrollY = (bottomValue: number) => {
  const scrollTop = window.scrollY || document.documentElement.scrollTop;
  return bottomValue + scrollTop;
};

export const union = (d1: IDOMRect, d2: IDOMRect): IDOMRect => {
  return {
    left: d1.left + d2.left,
    right: d1.right + d2.right,
    top: d1.top + d2.top,
    bottom: d1.bottom + d2.bottom,
    width: d1.width + d2.width,
    height: d1.height + d2.height,
  };
};

export const difference = (d1: IDOMRect, d2: IDOMRect): IDOMRect => {
  return {
    left: d1.left - d2.left,
    right: d1.right - d2.right,
    top: d1.top - d2.top,
    bottom: d1.bottom - d2.bottom,
    width: d1.width - d2.width,
    height: d1.height - d2.height,
  };
};

export const sum = (A: DOMRect[]): DOMRect => {
  if (A.length < 1) {
    throw new ReferenceError("vector A length is zero");
  }
  const last = A[A.length - 1];
  let x = 0;
  let y = 0;
  for (let i = 0; i < A.length; i++) {
    x += A[i].x;
    y += A[i].y;
  }
  return new DOMRect(x, y, last.width, last.height);
};

/**
 *
 * @param p1
 * @param childrens
 * @returns
 * @throws {Error} if childrens are empty
 */
export const positionComment = (
  p0Ref: React.RefObject<HTMLFormElement>,
  childrens: React.RefObject<HTMLLIElement>[]
): DOMRect => {
  if (!p0Ref.current) {
    throw new ReferenceError("p0Ref.curret is not defined");
  }
  const p0 = p0Ref.current.getBoundingClientRect();
  if (childrens.length > 0) {
    const childrensRect = childrens.map((child) => {
      if (!child.current) {
        throw new ReferenceError("child.current is not defined");
      }
      return child.current.getBoundingClientRect();
    });
    const childrensSum = sum(childrensRect);
    const total = childrens.length + 1;
    const x = childrensSum.x + p0.x;
    const y = childrensSum.y + p0.y;
    const averageX = x / total;
    const averageY = y / total; // + p0.height
    return new DOMRect(averageX, averageY, p0.width, p0.height);
  } else {
    return p0;
  }
};

export const commentsOverlapY = (
  commentsMap: Map<string, React.RefObject<HTMLFormElement>>,
  windowRef: React.RefObject<HTMLDivElement>
): Map<string, React.RefObject<HTMLFormElement>> => {
  const prevComments = Array.from(commentsMap.values());
  for (let i = 0; i < prevComments.length; i++) {
    for (let j = i + 1; j < prevComments.length; j++) {
      const vI = prevComments[i];
      const vJ = prevComments[j];
      if (vI.current && vJ.current) {
        const iRect = vI.current.getBoundingClientRect();
        const jRect = vJ.current.getBoundingClientRect();
        if (iRect && jRect) {
          const haveOverlap = iRect.bottom >= jRect.bottom;
          if (haveOverlap) {
            console.log("have overlap");
            vJ.current.style.top = `${
              offsetScrollY(iRect.top) + jRect.height
            }px`;
            vJ.current.style.bottom = `${
              limitScrollY(iRect.bottom) + jRect.height
            }px`;
          }
        }
      }
    }
  }
  adjustWindow(commentsMap, windowRef);
  return commentsMap;
};

export const adjustWindow = (
  commentsMap: Map<string, React.RefObject<HTMLFormElement>>,
  windowRef: React.RefObject<HTMLDivElement>
) => {
  const comments = Array.from(commentsMap.values());
  for (let i = 0; i < comments.length; i++) {
    const ref = comments[i];
    if (windowRef.current && ref.current) {
      const refRect = ref.current.getBoundingClientRect();
      const windowRect = windowRef.current?.getBoundingClientRect();
      if (windowRect.bottom < refRect.bottom) {
        const expand = refRect.bottom - windowRect.bottom + 20;
        windowRef.current.style.height = `${windowRect.height + expand}px`;
      }
    }
  }
};

export const addToMap = (
  map: Map<string, React.RefObject<HTMLFormElement>>,
  commentObject: CommentObject,
  value: React.RefObject<HTMLFormElement>
): Map<string, React.RefObject<any>> => {
  if (!value.current) {
    throw new ReferenceError("value.current not defiend in addToMap");
  }
  let p1 = positionComment(
    value,
    Array.from(commentObject.repliesRefs.values())
  );
  value.current.style.top = `${p1.top}px`;
  value.current.style.bottom = `${p1.bottom}px`;
  map.set(commentObject.color, value);
  const newMap = new Map(map);
  return newMap;
};

export const loadColorSquares = (
  colorSquares: ColorSquare[],
  colorsPicked: string[]
): ColorSquare[] => {
  let newSquares: ColorSquare[] = colorSquares;
  for (let i = 0; i < colorsPicked.length; i++) {
    newSquares = colorPicked(newSquares, colorsPicked[i]);
  }
  if (colorsPicked.length == 3) {
    newSquares = newSquares.map((square) => {
      return { ...square, fixed: true };
    });
  }
  return newSquares;
};

export const replyAlreadyExists = (
  comment: CommentObject,
  reply: ComplaintReply
): boolean => {
  const exists = comment.selectedReplies.find((r) => r.id === reply.id);
  return exists ? true : false;
};

/**
 *
 * @param comment
 * @param reply
 * @param replyRef
 * @returns
 * @throws {Error} if reply does not exist in comment
 */
export const removeReplyHighlight = (
  replyRef: React.RefObject<HTMLLIElement>
) => {
  replyRef!.current!.style.borderColor = "";
  replyRef!.current!.style.borderWidth = "0px";
  replyRef!.current!.style.borderRadius = "0px";
  return replyRef;
};

/**
 * Colors the reply ref, ads the reply to the comment and
 * updates the reply refs map of the comment
 * @param comment
 * @param reply
 * @param replyRef
 * @returns  newCommentObject[]
 * @throws {Error} if reply already exists in comment
 */
export const addReply = (
  comment: CommentObject,
  reply: ComplaintReply,
  replyRef: React.RefObject<HTMLLIElement>
): CommentObject => {
  const selectedReplies = comment.selectedReplies;
  const exists = replyAlreadyExists(comment, reply);
  if (exists) {
    throw new Error(
      `Reply ${reply.id} already exists in comment ${comment.color}`
    );
  }
  console.log(replyRef.current?.style);
  const styles = replyRef.current?.style;
  if (styles?.borderTopColor) {
    if (styles.borderTopColor != styles.borderBottomColor) {
      replyRef!.current!.style.borderRightColor = comment.color;
      replyRef!.current!.style.borderLeftColor = comment.color;
    } else {
      replyRef!.current!.style.borderBottomColor = comment.color;
    }
  } else {
    replyRef!.current!.style.borderColor = comment.color;
  }

  replyRef!.current!.style.borderWidth = "1px";
  replyRef!.current!.style.borderRadius = "5px";
  selectedReplies.push(reply);
  const refMap = new Map(comment.repliesRefs.entries());
  refMap.set(reply.id!, replyRef);
  return { ...comment, selectedReplies: selectedReplies, repliesRefs: refMap };
};

export const loadReplyReviews = (
  replyReviews: ReplyReview[],
  repliesMap: Map<string, React.RefObject<HTMLLIElement>>
): { loadedComments: CommentObject[]; colorsPicked: string[] } => {
  const loadedComments: CommentObject[] = [];
  const colorsPicked: string[] = [];
  for (let i = 0; i < replyReviews.length; i++) {
    const replyReview = replyReviews[i];
    let comment = createComment(replyReview.color, replyReview.reviewer);
    if (replyReview?.review?.comment != "") {
      comment = setBody(comment, replyReview?.review?.comment!);
      comment = setDone(comment, true);
    }
    for (let j = 0; j < replyReview.replies.length; j++) {
      const replyRef = repliesMap.get(replyReview?.replies[j]!.id!);
      if (!replyRef) {
        throw new Error(
          `replyRef ${replyReview?.replies[j]!.id} not found in repliesMap`
        );
      }

      if (!replyAlreadyExists(comment, replyReview?.replies[j]!)) {
        comment = addReply(comment, replyReview?.replies[j]!, replyRef);
      }
    }
    loadedComments.push(comment);
    colorsPicked.push(replyReview.color);
  }

  return { loadedComments, colorsPicked };
};

/**
 *
 * @param replyReviews
 * @param colorSquares
 * @param comments
 * @param repliesMap
 * @returns
 * @throws {Error} if replyRef not found in repliesMap
 * @throws {Error} if addReply throws an error because reply already exists in comment
 * @throws {Error} if addComment throws an error because comment already exists in comments
 * @throws {ColorAlreadyPickedError} if colorPicked throws an error because color already picked
 * @throws {Error} if colorPicked throws an error because color not found in colorSquares
 */
export const load = (
  user: Partial<User>,
  replyReviews: ReplyReview[],
  colorSquares: ColorSquare[],
  repliesMap: Map<string, React.RefObject<HTMLLIElement>>
): {
  newComments: CommentObject[];
  newSquares: ColorSquare[];
  actualColor: string;
} => {
  let actualColor = "";
  const { loadedComments, colorsPicked } = loadReplyReviews(
    replyReviews,
    repliesMap
  );
  let newSquares = loadColorSquares(colorSquares, colorsPicked);
  const unDoneComment = loadedComments.find(
    (comment) =>
      comment.reviewer.userName === user.userName && comment.comment == ""
  );
  if (unDoneComment) {
    newSquares = newSquares.map((square) =>
      square.color === unDoneComment.color
        ? { ...square, selected: true, fixed: true }
        : { ...square, fixed: true }
    );
    actualColor = unDoneComment.color;
  }
  return { newComments: loadedComments, newSquares: newSquares, actualColor };
};
