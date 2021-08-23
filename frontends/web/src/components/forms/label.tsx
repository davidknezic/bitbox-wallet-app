/**
 * Copyright 2018 Shift Devices AG
 * Copyright 2021 Shift Crypto AG
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { h, JSX } from 'preact';
import * as style from './label.css';

export function Label({
    className,
    children,
    id, // TODO: change to htmlFor when mirgated away from preact@8.x
    ...props
}: JSX.IntrinsicElements['label']) {
    const classes = [style.label, className].join(' ');
    return (
        <label for={id} className={classes} {...props}>
            {children}
        </label>
    );
}