// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

'use strict';

/**
 * `<mr-app>`
 *
 * The container component for all pages under the Monorail Polymer SPA.
 *
 */
class MrApp extends ReduxMixin(Polymer.Element) {
  static get is() {
    return 'mr-app';
  }

  static get properties() {
    return {
      loginUrl: String,
      logoutUrl: String,
      projectName: String,
      subheader: String,
      user: String,
      _boundLoadApprovalPage: {
        type: Function,
        value: function() {
          return this._loadApprovalPage.bind(this);
        },
      },
      formsToCheck: {
        type: Array,
        statePath: 'formsToCheck',
      },
      prevContext: {
        type: Object,
        statePath: 'prevContext',
      },
    };
  }

  connectedCallback() {
    super.connectedCallback();

    // TODO(zhangtiff): Figure out some way to save Redux state between
    //   page loads.

    page('*', (ctx, next) => {
      if (this.prevContext !== null) {
        // If this.prevContext is present, that means that we don't want to
        // navigate away from the last page, so we're restoring it.
        Object.assign(ctx, this.prevContext);
        this.dispatch({type: actionType.SET_CONTEXT, prevContext: null});
        // Set ctx.handled to false, so we don't push the state to browser's
        // history.
        ctx.handled = false;
        // We don't call next to avoid loading whatever page was supposed to
        // load next.
        return;
      }
      // Run query string parsing on all routes.
      // Based on: https://visionmedia.github.io/page.js/#plugins
      ctx.query = Qs.parse(ctx.querystring);

      next();
    });
    page('/p/:project/issues/approval', this._boundLoadApprovalPage);
    page.exit('*', (ctx, next) => {
      const isDirty = this.formsToCheck.some((form) =>
        (Object.keys(form.getDelta()).length !== 0));
      if (!isDirty || confirm('Discard your changes?')) {
        // Clear the forms to be checked, since we're navigating away.
        this.dispatch({type: actionType.CLEAR_FORMS_TO_CHECK});
      } else {
        // Store the current context, so that on the next page load we can
        // restore it.
        this.dispatch({type: actionType.SET_CONTEXT, prevContext: ctx});
      }
      next();
    });
    page();
  }

  loadWebComponent(name, props) {
    const component = document.createElement(name, {is: name});

    for (const key in props) {
      if (props.hasOwnProperty(key)) {
        component[key] = props[key];
      }
    }

    const main = Polymer.dom(this.root).querySelector('main');
    if (main) {
      // Clone the main tag without copying its children.
      const mainClone = main.cloneNode(false);
      mainClone.appendChild(component);

      main.parentNode.replaceChild(mainClone, main);
    }
  }

  _loadApprovalPage(ctx, next) {
    this.dispatch({
      type: actionType.UPDATE_ISSUE_REF,
      issueId: Number.parseInt(ctx.query.id),
      projectName: ctx.params.project,
    });

    this.projectName = ctx.params.project;
    this.subheader = 'Feature Launch Issue';

    this.loadWebComponent('mr-approval-page', {
      'projectName': ctx.params.project,
      'user': this.user,
    });
  }
}

customElements.define(MrApp.is, MrApp);
