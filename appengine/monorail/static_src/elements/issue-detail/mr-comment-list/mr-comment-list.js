// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

import '@polymer/polymer/polymer-legacy.js';
import {PolymerElement, html} from '@polymer/polymer';

import '../../chops/chops-button/chops-button.js';
import './mr-comment.js';
import {connectStore} from 'elements/reducers/base.js';
import * as issue from 'elements/reducers/issue.js';
import * as ui from 'elements/reducers/ui.js';
import 'elements/shared/mr-shared-styles.js';

const ADD_ISSUE_COMMENT_PERMISSION = 'addissuecomment';

/**
 * `<mr-comment-list>`
 *
 * Display a list of Monorail comments.
 *
 */
export class MrCommentList extends connectStore(PolymerElement) {
  static get template() {
    return html`
      <style include="mr-shared-styles">
        button.toggle {
          background: none;
          color: var(--chops-link-color);
          border: 0;
          border-bottom: var(--chops-normal-border);
          border-top: var(--chops-normal-border);
          width: 100%;
          padding: 0.5em 8px;
          text-align: left;
          font-size: var(--chops-main-font-size);
        }
        button.toggle:hover {
          cursor: pointer;
          text-decoration: underline;
        }
        button.toggle[hidden] {
          display: none;
        }
        .edit-slot {
          margin-top: 3em;
        }
      </style>
      <button on-click="toggleComments" class="toggle" hidden\$="[[_hideToggle]]">
        [[_computeCommentToggleVerb(_hideComments)]]
        [[_commentsHiddenCount]]
        older
        [[_pluralize(_commentsHiddenCount, 'comment')]]
      </button>
      <template is="dom-repeat" items="[[comments]]" as="comment">
        <template is="dom-if" if="[[_computeShowComment(_hideComments, _commentsHiddenCount, index)]]">
          <mr-comment
            comment="[[comment]]"
            heading-level="[[headingLevel]]"
            highlighted="[[_computeHighlighted(focusId, comment)]]"
            quick-mode="[[quickMode]]"
          ></mr-comment>
        </template>
      </template>
      <div class="edit-slot" hidden$="[[!_canAddComment(issuePermissions)]]">
        <slot></slot>
      </div>
    `;
  }

  static get is() {
    return 'mr-comment-list';
  }

  static get properties() {
    return {
      commentsShownCount: {
        type: Number,
        value: 2,
      },
      comments: {
        type: Array,
        value: [],
      },
      headingLevel: {
        type: Number,
        value: 4,
      },
      issuePermissions: Object,
      focusId: String,
      quickMode: {
        type: Boolean,
        value: false,
      },
      _hideComments: {
        type: Boolean,
        value: true,
      },
      _commentsHiddenCount: {
        type: Number,
        computed: '_computeCommentsHiddenCount(commentsShownCount, comments.length)',
      },
      _hideToggle: {
        type: Boolean,
        value: false,
        computed: '_computeHideToggle(_commentsHiddenCount)',
      },
    };
  }

  static get observers() {
    return [
      '_onFocusId(' +
        'focusId, comments, _hideComments, _commentsHiddenCount)',
    ];
  }

  stateChanged(state) {
    this.setProperties({
      focusId: ui.focusId(state),
      issuePermissions: issue.permissions(state),
    });
  }

  toggleComments() {
    this._hideComments = !this._hideComments;
  }

  _canAddComment(issuePermissions) {
    return (issuePermissions || []).includes(ADD_ISSUE_COMMENT_PERMISSION);
  }

  _computeShowComment(hideComments, commentsHiddenCount, index) {
    return !hideComments || index >= commentsHiddenCount;
  }

  _computeHighlighted(focusId, comment) {
    if (!comment.sequenceNum || !focusId) return;
    return 'c' + comment.sequenceNum === focusId;
  }

  _computeCommentsHiddenCount(shownCount, numComments) {
    return Math.max(numComments - shownCount, 0);
  }

  _computeHideToggle(hiddenCount) {
    return hiddenCount <= 0;
  }

  _computeCommentToggleVerb(hideComments) {
    return hideComments ? 'Show' : 'Hide';
  }

  _pluralize(count, baseWord, pluralWord) {
    pluralWord = pluralWord || `${baseWord}s`;
    return count == 1 ? baseWord : pluralWord;
  }

  _onFocusId(focusId, comments, hideComments, commentsHiddenCount) {
    if (!comments || !focusId || !hideComments) return;
    for (let i = 0; i < comments.length; ++i) {
      const commentId = 'c' + comments[i].sequenceNum;
      if (commentId === focusId &&
          !this._computeShowComment(hideComments, commentsHiddenCount, i)) {
        this._hideComments = false;
        break;
      }
    }
  }
}
customElements.define(MrCommentList.is, MrCommentList);
