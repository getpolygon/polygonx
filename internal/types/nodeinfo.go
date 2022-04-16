// BSD 3-Clause License

// Copyright (c) 2021, Michael Grigoryan
// All rights reserved.

// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:

// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.

// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.

// 3. Neither the name of the copyright holder nor the names of its
//    contributors may be used to endorse or promote products derived from
//    this software without specific prior written permission.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
package types

// Empty type for complying with the specification in NodeInfo
type NodeInfoService string

// Empty type for complying with the specification in NodeInfo
type NodeInfoProtocol string

type NodeInfoSoftware struct {
	// The canonical name of this server software.
	Name string `json:"name"`
	// The version of this server software.
	Version string `json:"version"`
	// The url of the source code repository of this server software.
	Homepage string `json:"homepage"`
	// The url of the homepage of this server software.
	Repository string `json:"repository"`
}

type NodeInfoUsers struct {
	// The total amount of on this server registered users.
	Total int64 `json:"total"`
	// The amount of users that signed in at least once in the last 180 days.
	ActiveHalfYear int64 `json:"activeHalfYear"`
	// The amount of users that signed in at least once in the last 30 days.
	ActiveOneMonth int64 `json:"activeOneMonth"`
}

type NodeInfoUsage struct {
	// Usage statistics for this server.
	Users NodeInfoUsers `json:"users"`
	// The amount of posts that were made by users that are registered on this server.
	LocalPosts int64 `json:"localPosts"`
	// The amount of comments that were made by users that are registered on this server.
	LocalComments int64 `json:"localComments"`
}

// NodeInfo is an effort to create a standardized way of exposing metadata about
// a server running one of the distributed social networks. The two key goals are
// being able to get better insights into the user base of distributed social networking
// and the ability to build tools that allow users to choose the best fitting software
// and server for their needs. Polygon uses version 2.1 of the NodeInfo specification.
// http://nodeinfo.diaspora.software/docson/index.html#/ns/schema/2.1#$$expand
type NodeInfo struct {
	// The schema version, must be 2.1.
	Version string `json:"version"`

	// Usage statistics for this server
	Usage NodeInfoUsage `json:"usage"`

	// Metadata about server software in use.
	Software NodeInfoSoftware `json:"software"`

	// The protocols supported on this server.
	Protocols []NodeInfoProtocol `json:"protocols"`

	// The third party sites this server can connect to via their application API.
	Services []NodeInfoService `json:"services"`

	// Whether this server allows open self-registration.
	OpenRegistrations bool `json:"openRegistrations"`

	// Free form key value pairs for software specific values. Clients should not rely on any specific key present.
	Metadata map[string]interface{} `json:"metadata"`
}
