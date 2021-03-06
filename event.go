// Copyright 2016 the Go-Circuits Authors.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

func NewEvent(channel, target string) Event {
	return &BaseEvent{channel: channel, target: target}
}

type Event interface {
	Channel() string
	Target() string
	SetNotifyFailure(bool)
	NotifyFailure() bool
	SetNotifySuccess(bool)
	NotifySuccess() bool
	SetNotifyComplete(bool)
	NotifyComplete() bool
}

type BaseEvent struct {
	channel         string
	target          string
	notify_failure  bool
	notify_success  bool
	notify_complete bool
}

func (e *BaseEvent) Channel() string {
	return e.channel
}

func (e *BaseEvent) Target() string {
	return e.target
}

func (e *BaseEvent) SetNotifyFailure(v bool) {
	e.notify_failure = v
}

func (e *BaseEvent) NotifyFailure() bool {
	return e.notify_failure
}

func (e *BaseEvent) SetNotifySuccess(v bool) {
	e.notify_success = v
}

func (e *BaseEvent) NotifySuccess() bool {
	return e.notify_success
}

func (e *BaseEvent) SetNotifyComplete(v bool) {
	e.notify_complete = v
}

func (e *BaseEvent) NotifyComplete() bool {
	return e.notify_complete
}
