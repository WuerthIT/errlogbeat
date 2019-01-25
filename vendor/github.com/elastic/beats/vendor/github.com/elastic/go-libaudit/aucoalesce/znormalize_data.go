// mknormalize_data.go
// MACHINE GENERATED BY THE ABOVE COMMAND; DO NOT EDIT

// Copyright 2017-2018 Elasticsearch Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aucoalesce

import (
	"encoding/base64"
	"fmt"
)

var assets map[string][]byte

func asset(key string) ([]byte, error) {
	if assets == nil {
		assets = map[string][]byte{}

		var value []byte
		value, _ = base64.StdEncoding.DecodeString("LS0tCiMgTWFjcm9zIGRlY2xhcmVzIHNvbWUgWUFNTCBhbmNob3JzIHRoYXQgY2FuIGJlIHJlZmVyZW5jZWQgZm9yIHNvbWUgY29tbW9uCiMgb2JqZWN0IHR5cGUgbm9ybWFsaXphdGlvbnMgbGlrZSB1c2VyLXNlc3Npb24sIHNvY2tldCwgb3IgcHJvY2Vzcy4KbWFjcm9zOgotICZkZWZhdWx0cwogIHN1YmplY3Q6CiAgICBwcmltYXJ5OiBhdWlkCiAgICBzZWNvbmRhcnk6IHVpZAogIGhvdzogW2V4ZSwgY29tbV0KCi0gJm1hY3JvLXVzZXItc2Vzc2lvbgogIHN1YmplY3Q6CiAgICBwcmltYXJ5OiBhdWlkCiAgICBzZWNvbmRhcnk6IFthY2N0LCBpZCwgdWlkXQogIG9iamVjdDoKICAgIHByaW1hcnk6IHRlcm1pbmFsCiAgICBzZWNvbmRhcnk6IFthZGRyLCBob3N0bmFtZV0KICAgIHdoYXQ6IHVzZXItc2Vzc2lvbgogIGhvdzogW2V4ZSwgdGVybWluYWxdCgotICZtYWNyby1zb2NrZXQKICA8PDogKmRlZmF1bHRzCiAgb2JqZWN0OgogICAgcHJpbWFyeTogW2FkZHIsIHBhdGhdCiAgICBzZWNvbmRhcnk6IHBvcnQKICAgIHdoYXQ6IHNvY2tldAoKLSAmbWFjcm8tcHJvY2VzcwogIDw8OiAqZGVmYXVsdHMKICBvYmplY3Q6CiAgICBwcmltYXJ5OiBbY21kLCBleGUsIGNvbW1dCiAgICBzZWNvbmRhcnk6IHBpZAogICAgd2hhdDogcHJvY2VzcwogIGhvdzogdGVybWluYWwKCiMgTm9ybWFsaXphdGlvbnMgaXMgYSBsaXN0IG9mIGRlY2xhcmF0aW9ucyBzcGVjaWZ5aW5nIGhvdyB0byBub3JtYWxpemUgdGhlIGRhdGEKIyBjb250YWluZWQgaW4gYW4gZXZlbnQuIFRoZSBub3JtYWxpemF0aW9uIGNhbiBiZSBhcHBsaWVkIGJhc2VkIG9uIHRoZSBzeXNjYWxsCiMgbmFtZSAoZS5nLiBjb25uZWN0LCBvcGVuKSBvciBiYXNlZCBvbiB0aGUgcmVjb3JkIHR5cGUgKGUuZy4gVVNFUl9MT0dJTikuCiMgTm8gdHdvIG5vcm1hbGl6YXRpb25zIGNhbiBhcHBseSB0byB0aGUgc2FtZSBzeXNjYWxsIG9yIHJlY29yZCB0eXBlLiBUaGlzCiMgd2lsbCByZXN1bHQgaW4gYSBmYWlsdXJlIGF0IGxvYWQgdGltZS4KIwojIEVhY2ggbm9ybWFsaXphdGlvbiBzaG91bGQgc3BlY2lmeToKIyAgIGFjdGlvbiAtIHdoYXQgaGFwcGVuZWQKIyAgIGFjdG9yICAtIHdobyBkaWQgdGhpcyBvciB3aG8gdHJpZ2dlcmVkIHRoZSBldmVudAojICAgb2JqZWN0IC0gd2hhdCB3YXMgdGhlICJ0aGluZyIgaW52b2x2ZWQgaW4gdGhlIGFjdGlvbiAoZS5nLiBwcm9jZXNzLCBzb2NrZXQpCiMgICBob3cgICAgLSBob3cgd2FzIHRoZSBhY3Rpb24gcGVyZm9ybWVkIChlLmcuIGV4ZSBvciB0ZXJtaW5hbCkKbm9ybWFsaXphdGlvbnM6Ci0KICBhY3Rpb246IG9wZW5lZC1maWxlCiAgb2JqZWN0OgogICAgd2hhdDogZmlsZQogIHN5c2NhbGxzOgogIC0gY3JlYXQKICAtIGZhbGxvY2F0ZQogIC0gdHJ1bmNhdGUKICAtIGZ0cnVuY2F0ZQogIC0gb3BlbgogIC0gb3BlbmF0CiAgLSByZWFkbGluawogIC0gcmVhZGxpbmthdAotCiAgYWN0aW9uOiBjaGFuZ2VkLWZpbGUtYXR0cmlidXRlcy1vZgogIG9iamVjdDoKICAgIHdoYXQ6IGZpbGUKICBzeXNjYWxsczoKICAtIHNldHhhdHRyCiAgLSBmc2V0eGF0dHIKICAtIGxzZXR4YXR0cgogIC0gcmVtb3ZleGF0dHIKICAtIGZyZW1vdmV4YXR0cgogIC0gbHJlbW92ZXhhdHRyCi0KICBhY3Rpb246IGNoYW5nZWQtZmlsZS1wZXJtaXNzaW9ucy1vZgogIG9iamVjdDoKICAgIHdoYXQ6IGZpbGUKICBzeXNjYWxsczoKICAtIGNobW9kCiAgLSBmY2htb2QKICAtIGZjaG1vZGF0Ci0KICBhY3Rpb246IGNoYW5nZWQtZmlsZS1vd25lcnNoaXAtb2YKICBvYmplY3Q6CiAgICB3aGF0OiBmaWxlCiAgc3lzY2FsbHM6CiAgLSBjaG93bgogIC0gZmNob3duCiAgLSBmY2hvd25hdAogIC0gbGNob3duCi0KICBhY3Rpb246IGxvYWRlZC1rZXJuZWwtbW9kdWxlCiAgb2JqZWN0OgogICAgd2hhdDogZmlsZQogICAgcHJpbWFyeTogbmFtZQogIHJlY29yZF90eXBlczoKICAtIEtFUk5fTU9EVUxFCiAgc3lzY2FsbHM6CiAgLSBmaW5pdF9tb2R1bGUKICAtIGluaXRfbW9kdWxlCi0KICBhY3Rpb246IHVubG9hZGVkLWtlcm5lbC1tb2R1bGUKICBvYmplY3Q6CiAgICB3aGF0OiBmaWxlCiAgc3lzY2FsbHM6CiAgLSBkZWxldGVfbW9kdWxlCi0KICBhY3Rpb246IGNyZWF0ZWQtZGlyZWN0b3J5CiAgb2JqZWN0OgogICAgd2hhdDogZmlsZQogICAgcGF0aF9pbmRleDogMQogIHN5c2NhbGxzOgogIC0gbWtkaXIKICAtIG1rZGlyYXQKLQogIGFjdGlvbjogbW91bnRlZAogIG9iamVjdDoKICAgIHdoYXQ6IGZpbGVzeXN0ZW0KICAgIHBhdGhfaW5kZXg6IDEKICBzeXNjYWxsczoKICAtIG1vdW50Ci0KICBhY3Rpb246IHJlbmFtZWQKICBvYmplY3Q6CiAgICB3aGF0OiBmaWxlCiAgICBwYXRoX2luZGV4OiAyCiAgc3lzY2FsbHM6CiAgLSByZW5hbWUKICAtIHJlbmFtZWF0CiAgLSByZW5hbWVhdDIKLQogIGFjdGlvbjogY2hlY2tlZC1tZXRhZGF0YS1vZgogIG9iamVjdDoKICAgIHdoYXQ6IGZpbGUKICBzeXNjYWxsczoKICAtIGFjY2VzcwogIC0gZmFjY2Vzc2F0CiAgLSBuZXdmc3RhdGF0CiAgLSBzdGF0CiAgLSBmc3RhdAogIC0gbHN0YXQKICAtIHN0YXQ2NAogIC0gZ2V0eGF0dHIKICAtIGxnZXR4YXR0cgogIC0gZmdldHhhdHRyCi0KICBhY3Rpb246IGNoZWNrZWQtZmlsZXN5c3RlbS1tZXRhZGF0YS1vZgogIG9iamVjdDoKICAgIHdoYXQ6IGZpbGVzeXN0ZW0KICBzeXNjYWxsczoKICAtIHN0YXRmcwogIC0gZnN0YXRmcwotCiAgYWN0aW9uOiBzeW1saW5rZWQKICBvYmplY3Q6CiAgICB3aGF0OiBmaWxlCiAgc3lzY2FsbHM6CiAgLSBzeW1saW5rCiAgLSBzeW1saW5rYXQKLQogIGFjdGlvbjogdW5tb3VudGVkCiAgb2JqZWN0OgogICAgd2hhdDogZmlsZXN5c3RlbQogIHN5c2NhbGxzOgogIC0gdW1vdW50MgotCiAgYWN0aW9uOiBkZWxldGVkCiAgb2JqZWN0OgogICAgd2hhdDogZmlsZQogIHN5c2NhbGxzOgogIC0gcm1kaXIKICAtIHVubGluawogIC0gdW5saW5rYXQKLQogIGFjdGlvbjogY2hhbmdlZC10aW1lc3RhbXAtb2YKICBvYmplY3Q6CiAgICB3aGF0OiBmaWxlCiAgc3lzY2FsbHM6CiAgLSB1dGltZQogIC0gdXRpbWVzCiAgLSBmdXRpbWVzYXQKICAtIGZ1dGltZW5zCiAgLSB1dGltZW5zYXQKLQogIGFjdGlvbjogZXhlY3V0ZWQKICBvYmplY3Q6CiAgICB3aGF0OiBmaWxlCiAgc3lzY2FsbHM6CiAgLSBleGVjdmUKICAtIGV4ZWN2ZWF0Ci0KICBhY3Rpb246IGxpc3Rlbi1mb3ItY29ubmVjdGlvbnMKICBvYmplY3Q6CiAgICB3aGF0OiBzb2NrZXQKICBzeXNjYWxsczoKICAtIGxpc3RlbgotCiAgYWN0aW9uOiBhY2NlcHRlZC1jb25uZWN0aW9uLWZyb20KICBvYmplY3Q6CiAgICB3aGF0OiBzb2NrZXQKICBzeXNjYWxsczoKICAtIGFjY2VwdAogIC0gYWNjZXB0NAotCiAgYWN0aW9uOiBib3VuZC1zb2NrZXQKICBvYmplY3Q6CiAgICB3aGF0OiBzb2NrZXQKICBzeXNjYWxsczoKICAtIGJpbmQKLQogIGFjdGlvbjogY29ubmVjdGVkLXRvCiAgb2JqZWN0OgogICAgd2hhdDogc29ja2V0CiAgc3lzY2FsbHM6CiAgLSBjb25uZWN0Ci0KICBhY3Rpb246IHJlY2VpdmVkLWZyb20KICBvYmplY3Q6CiAgICB3aGF0OiBzb2NrZXQKICBzeXNjYWxsczoKICAtIHJlY3Zmcm9tCiAgLSByZWN2bXNnCi0KICBhY3Rpb246IHNlbnQtdG8KICBvYmplY3Q6CiAgICB3aGF0OiBzb2NrZXQKICBzeXNjYWxsczoKICAtIHNlbmR0bwogIC0gc2VuZG1zZwotCiAgYWN0aW9uOiBraWxsZWQtcGlkCiAgb2JqZWN0OgogICAgd2hhdDogcHJvY2VzcwogIHN5c2NhbGxzOgogIC0ga2lsbAogIC0gdGtpbGwKICAtIHRna2lsbAotCiAgYWN0aW9uOiBjaGFuZ2VkLWlkZW50aXR5LW9mCiAgb2JqZWN0OgogICAgd2hhdDogcHJvY2VzcwogIGhvdzogc3lzY2FsbAogIHN5c2NhbGxzOgogIC0gc2V0dWlkCiAgLSBzZXRldWlkCiAgLSBzZXRmc3VpZAogIC0gc2V0cmV1aWQKICAtIHNldHJlc3VpZAogIC0gc2V0Z2lkCiAgLSBzZXRlZ2lkCiAgLSBzZXRmc2dpZAogIC0gc2V0cmVnaWQKICAtIHNldHJlc2dpZAotCiAgYWN0aW9uOiBjaGFuZ2VkLXN5c3RlbS10aW1lCiAgb2JqZWN0OgogICAgd2hhdDogc3lzdGVtCiAgc3lzY2FsbHM6CiAgLSBzZXR0aW1lb2ZkYXkKICAtIGNsb2NrX3NldHRpbWUKICAtIHN0aW1lCiAgLSBhZGp0aW1leAotCiAgYWN0aW9uOiBtYWtlLWRldmljZQogIG9iamVjdDoKICAgIHdoYXQ6IGZpbGUKICBzeXNjYWxsczoKICAtIG1rbm9kCiAgLSBta25vZGF0Ci0KICBhY3Rpb246IGNoYW5nZWQtc3lzdGVtLW5hbWUKICBvYmplY3Q6CiAgICB3aGF0OiBzeXN0ZW0KICBzeXNjYWxsczoKICAtIHNldGhvc3RuYW1lCiAgLSBzZXRkb21haW5uYW1lCi0KICBhY3Rpb246IGFsbG9jYXRlZC1tZW1vcnkKICBvYmplY3Q6CiAgICB3aGF0OiBtZW1vcnkKICBzeXNjYWxsczoKICAtIG1tYXAKICAtIGJyawotCiAgYWN0aW9uOiBhZGp1c3RlZC1zY2hlZHVsaW5nLXBvbGljeS1vZgogIG9iamVjdDoKICAgIHdoYXQ6IHByb2Nlc3MKICBob3c6IHN5c2NhbGwKICBzeXNjYWxsczoKICAtIHNjaGVkX3NldHBhcmFtCiAgLSBzY2hlZF9zZXRzY2hlZHVsZXIKICAtIHNjaGVkX3NldGF0dHIKLQogIGFjdGlvbjogY2F1c2VkLW1hYy1wb2xpY3ktZXJyb3IKICBvYmplY3Q6CiAgICB3aGF0OiBzeXN0ZW0KICByZWNvcmRfdHlwZXM6IFNFTElOVVhfRVJSCi0KICBhY3Rpb246IGxvYWRlZC1maXJld2FsbC1ydWxlLXRvCiAgb2JqZWN0OgogICAgcHJpbWFyeTogdGFibGUKICAgIHdoYXQ6IGZpcmV3YWxsCiAgcmVjb3JkX3R5cGVzOiBORVRGSUxURVJfQ0ZHCi0KICAjIENvdWxkIGJlIGVudGVyZWQgb3IgZXhpdGVkIGJhc2VkIG9uIHByb20gZmllbGQuCiAgYWN0aW9uOiBjaGFuZ2VkLXByb21pc2N1b3VzLW1vZGUtb24tZGV2aWNlCiAgb2JqZWN0OgogICAgcHJpbWFyeTogZGV2CiAgICB3aGF0OiBuZXR3b3JrLWRldmljZQogIHJlY29yZF90eXBlczogQU5PTV9QUk9NSVNDVU9VUwotCiAgYWN0aW9uOiBsb2NrZWQtYWNjb3VudAogIHJlY29yZF90eXBlczogQUNDVF9MT0NLCi0KICBhY3Rpb246IHVubG9ja2VkLWFjY291bnQKICByZWNvcmRfdHlwZXM6IEFDQ1RfVU5MT0NLCi0KICBhY3Rpb246IGFkZGVkLWdyb3VwLWFjY291bnQtdG8KICBvYmplY3Q6CiAgICBwcmltYXJ5OiBbaWQsIGFjY3RdCiAgICB3aGF0OiBhY2NvdW50CiAgcmVjb3JkX3R5cGVzOiBBRERfR1JPVVAKLQogIGFjdGlvbjogYWRkZWQtdXNlci1hY2NvdW50CiAgb2JqZWN0OgogICAgcHJpbWFyeTogW2lkLCBhY2N0XQogICAgd2hhdDogYWNjb3VudAogIHJlY29yZF90eXBlczogQUREX1VTRVIKLQogIGFjdGlvbjogY3Jhc2hlZC1wcm9ncmFtCiAgb2JqZWN0OgogICAgcHJpbWFyeTogW2NvbW0sIGV4ZV0KICAgIHNlY29uZGFyeTogcGlkCiAgICB3aGF0OiBwcm9jZXNzCiAgaG93OiBzaWcKICByZWNvcmRfdHlwZXM6IEFOT01fQUJFTkQKLQogIGFjdGlvbjogYXR0ZW1wdGVkLWV4ZWN1dGlvbi1vZi1mb3JiaWRkZW4tcHJvZ3JhbQogIG9iamVjdDoKICAgIHByaW1hcnk6IGNtZAogICAgd2hhdDogcHJvY2VzcwogIGhvdzogdGVybWluYWwKICByZWNvcmRfdHlwZXM6IEFOT01fRVhFQwotCiAgYWN0aW9uOiB1c2VkLXN1c3BjaW91cy1saW5rCiAgcmVjb3JkX3R5cGVzOiBBTk9NX0xJTksKLQogIDw8OiAqbWFjcm8tdXNlci1zZXNzaW9uCiAgYWN0aW9uOiBmYWlsZWQtbG9nLWluLXRvby1tYW55LXRpbWVzLXRvCiAgcmVjb3JkX3R5cGVzOiBBTk9NX0xPR0lOX0ZBSUxVUkVTCi0KICA8PDogKm1hY3JvLXVzZXItc2Vzc2lvbgogIGFjdGlvbjogYXR0ZW1wdGVkLWxvZy1pbi1mcm9tLXVudXN1YWwtcGxhY2UtdG8KICByZWNvcmRfdHlwZXM6IEFOT01fTE9HSU5fTE9DQVRJT04KLQogIDw8OiAqbWFjcm8tdXNlci1zZXNzaW9uCiAgYWN0aW9uOiBvcGVuZWQtdG9vLW1hbnktc2Vzc2lvbnMtdG8KICByZWNvcmRfdHlwZXM6IEFOT01fTE9HSU5fU0VTU0lPTlMKLQogIDw8OiAqbWFjcm8tdXNlci1zZXNzaW9uCiAgYWN0aW9uOiBhdHRlbXB0ZWQtbG9nLWluLWR1cmluZy11bnVzdWFsLWhvdXItdG8KICByZWNvcmRfdHlwZXM6IEFOT01fTE9HSU5fVElNRQotCiAgYWN0aW9uOiB0ZXN0ZWQtZmlsZS1zeXN0ZW0taW50ZWdyaXR5LW9mCiAgb2JqZWN0OgogICAgcHJpbWFyeTogaG9zdG5hbWUKICAgIHdoYXQ6IGZpbGVzeXN0ZW0KICByZWNvcmRfdHlwZXM6IEFOT01fUkJBQ19JTlRFR1JJVFlfRkFJTAotCiAgYWN0aW9uOiB2aW9sYXRlZC1zZWxpbnV4LXBvbGljeQogIHN1YmplY3Q6CiAgICBwcmltYXJ5OiBzY29udGV4dAogIG9iamVjdDoKICAgIHByaW1hcnk6IHRjb250ZXh0CiAgcmVjb3JkX3R5cGVzOiBBVkMKLQogIGFjdGlvbjogY2hhbmdlZC1ncm91cAogIHJlY29yZF90eXBlczogQ0hHUlBfSUQKLQogIGFjdGlvbjogY2hhbmdlZC11c2VyLWlkCiAgcmVjb3JkX3R5cGVzOiBDSFVTRVJfSUQKLQogIGFjdGlvbjogY2hhbmdlZC1hdWRpdC1jb25maWd1cmF0aW9uCiAgb2JqZWN0OgogICAgcHJpbWFyeTogW29wLCBrZXksIGF1ZGl0X2VuYWJsZWQsIGF1ZGl0X3BpZCwgYXVkaXRfYmFja2xvZ19saW1pdCwgYXVkaXRfZmFpbHVyZV0KICAgIHdoYXQ6IGF1ZGl0LWNvbmZpZwogIHJlY29yZF90eXBlczogQ09ORklHX0NIQU5HRQotCiAgPDw6ICptYWNyby11c2VyLXNlc3Npb24KICBhY3Rpb246IGFjcXVpcmVkLWNyZWRlbnRpYWxzCiAgcmVjb3JkX3R5cGVzOiBDUkVEX0FDUQotCiAgPDw6ICptYWNyby11c2VyLXNlc3Npb24KICBhY3Rpb246IGRpc3Bvc2VkLWNyZWRlbnRpYWxzCiAgcmVjb3JkX3R5cGVzOiBDUkVEX0RJU1AKLQogIDw8OiAqbWFjcm8tdXNlci1zZXNzaW9uCiAgYWN0aW9uOiByZWZyZXNoZWQtY3JlZGVudGlhbHMKICByZWNvcmRfdHlwZXM6IENSRURfUkVGUgotCiAgPDw6ICptYWNyby11c2VyLXNlc3Npb24KICBhY3Rpb246IG5lZ290aWF0ZWQtY3J5cHRvLWtleQogIG9iamVjdDoKICAgIHByaW1hcnk6IGZwCiAgICBzZWNvbmRhcnk6IFthZGRyLCBob3N0bmFtZV0KICAgIHdoYXQ6IHVzZXItc2Vzc2lvbgogIHJlY29yZF90eXBlczogQ1JZUFRPX0tFWV9VU0VSCiAgc291cmNlX2lwOiBbYWRkcl0KLQogIGFjdGlvbjogY3J5cHRvLW9mZmljZXItbG9nZ2VkLWluCiAgcmVjb3JkX3R5cGVzOiBDUllQVE9fTE9HSU4KLQogIGFjdGlvbjogY3J5cHRvLW9mZmljZXItbG9nZ2VkLW91dAogIHJlY29yZF90eXBlczogQ1JZUFRPX0xPR09VVAotCiAgPDw6ICptYWNyby11c2VyLXNlc3Npb24KICBhY3Rpb246IHN0YXJ0ZWQtY3J5cHRvLXNlc3Npb24KICBvYmplY3Q6CiAgICBwcmltYXJ5OiBhZGRyCiAgICBzZWNvbmRhcnk6IFtycG9ydF0KICByZWNvcmRfdHlwZXM6IENSWVBUT19TRVNTSU9OCiAgc291cmNlX2lwOiBbYWRkcl0KLQogIGFjdGlvbjogYWNjZXNzLXJlc3VsdAogIHJlY29yZF90eXBlczogREFDX0NIRUNLCi0KICBhY3Rpb246IGFib3J0ZWQtYXVkaXRkLXN0YXJ0dXAKICBvYmplY3Q6CiAgICB3aGF0OiBzZXJ2aWNlCiAgcmVjb3JkX3R5cGVzOiBEQUVNT05fQUJPUlQKLQogIGFjdGlvbjogcmVtb3RlLWF1ZGl0LWNvbm5lY3RlZAogIG9iamVjdDoKICAgIHdoYXQ6IHNlcnZpY2UKICByZWNvcmRfdHlwZXM6IERBRU1PTl9BQ0NFUFQKLQogIGFjdGlvbjogcmVtb3RlLWF1ZGl0LWRpc2Nvbm5lY3RlZAogIG9iamVjdDoKICAgIHdoYXQ6IHNlcnZpY2UKICByZWNvcmRfdHlwZXM6IERBRU1PTl9DTE9TRQotCiAgYWN0aW9uOiBjaGFuZ2VkLWF1ZGl0ZC1jb25maWd1cmF0aW9uCiAgb2JqZWN0OgogICAgd2hhdDogc2VydmljZQogIHJlY29yZF90eXBlczogREFFTU9OX0NPTkZJRwotCiAgYWN0aW9uOiBzaHV0ZG93bi1hdWRpdAogIG9iamVjdDoKICAgIHdoYXQ6IHNlcnZpY2UKICByZWNvcmRfdHlwZXM6IERBRU1PTl9FTkQKLQogIGFjdGlvbjogYXVkaXQtZXJyb3IKICBvYmplY3Q6CiAgICB3aGF0OiBzZXJ2aWNlCiAgcmVjb3JkX3R5cGVzOiBEQUVNT05fRVJSCi0KICBhY3Rpb246IHJlY29uZmlndXJlZC1hdWRpdGQKICBvYmplY3Q6CiAgICB3aGF0OiBzZXJ2aWNlCiAgcmVjb3JkX3R5cGVzOiBEQUVNT05fUkVDT05GSUcKLQogIGFjdGlvbjogcmVzdW1lZC1hdWRpdC1sb2dnaW5nCiAgb2JqZWN0OgogICAgd2hhdDogc2VydmljZQogIHJlY29yZF90eXBlczogREFFTU9OX1JFU1VNRQotCiAgYWN0aW9uOiByb3RhdGVkLWF1ZGl0LWxvZ3MKICBvYmplY3Q6CiAgICB3aGF0OiBzZXJ2aWNlCiAgcmVjb3JkX3R5cGVzOiBEQUVNT05fUk9UQVRFCi0KICBhY3Rpb246IHN0YXJ0ZWQtYXVkaXQKICBvYmplY3Q6CiAgICB3aGF0OiBzZXJ2aWNlCiAgcmVjb3JkX3R5cGVzOiBEQUVNT05fU1RBUlQKLQogIGFjdGlvbjogZGVsZXRlZC1ncm91cC1hY2NvdW50LWZyb20KICBvYmplY3Q6CiAgICBwcmltYXJ5OiBbaWQsIGFjY3RdCiAgICB3aGF0OiBhY2NvdW50CiAgcmVjb3JkX3R5cGVzOiBERUxfR1JPVVAKLQogIGFjdGlvbjogZGVsZXRlZC11c2VyLWFjY291bnQKICBvYmplY3Q6CiAgICBwcmltYXJ5OiBbaWQsIGFjY3RdCiAgICB3aGF0OiBhY2NvdW50CiAgcmVjb3JkX3R5cGVzOiBERUxfVVNFUgotCiAgYWN0aW9uOiBjaGFuZ2VkLWF1ZGl0LWZlYXR1cmUKICBvYmplY3Q6CiAgICBwcmltYXJ5OiBmZWF0dXJlCiAgICB3aGF0OiBzeXN0ZW0KICByZWNvcmRfdHlwZXM6IEZFQVRVUkVfQ0hBTkdFCi0KICBhY3Rpb246IHJlbGFiZWxlZC1maWxlc3lzdGVtCiAgcmVjb3JkX3R5cGVzOiBGU19SRUxBQkVMCi0KICBhY3Rpb246IGF1dGhlbnRpY2F0ZWQtdG8tZ3JvdXAKICByZWNvcmRfdHlwZXM6IEdSUF9BVVRICi0KICA8PDogKm1hY3JvLXVzZXItc2Vzc2lvbgogIGFjdGlvbjogY2hhbmdlZC1ncm91cC1wYXNzd29yZAogIG9iamVjdDoKICAgIHByaW1hcnk6IGFjY3QKICAgIHdoYXQ6IHVzZXItc2Vzc2lvbgogIHJlY29yZF90eXBlczogR1JQX0NIQVVUSFRPSwotCiAgYWN0aW9uOiBtb2RpZmllZC1ncm91cC1hY2NvdW50CiAgb2JqZWN0OgogICAgcHJpbWFyeTogW2lkLCBhY2N0XQogICAgd2hhdDogYWNjb3VudAogIHJlY29yZF90eXBlczogR1JQX01HTVQKLQogIGFjdGlvbjogaW5pdGlhbGl6ZWQtYXVkaXQtc3Vic3lzdGVtCiAgcmVjb3JkX3R5cGVzOiBLRVJORUwKLQogIGFjdGlvbjogbW9kaWZpZWQtbGV2ZWwtb2YKICBvYmplY3Q6CiAgICBwcmltYXJ5OiBwcmludGVyCiAgICB3aGF0OiBwcmludGVyCiAgcmVjb3JkX3R5cGVzOiBMQUJFTF9MRVZFTF9DSEFOR0UKLQogIGFjdGlvbjogb3ZlcnJvZGUtbGFiZWwtb2YKICBvYmplY3Q6CiAgICB3aGF0OiBtYWMtY29uZmlnCiAgcmVjb3JkX3R5cGVzOiBMQUJFTF9PVkVSUklERQotCiAgb2JqZWN0OgogICAgd2hhdDogbWFjLWNvbmZpZwogIHJlY29yZF90eXBlczoKICAtIEFVRElUX0RFVl9BTExPQwogIC0gQVVESVRfREVWX0RFQUxMT0MKICAtIEFVRElUX0ZTX1JFTEFCRUwKICAtIEFVRElUX1VTRVJfTUFDX1BPTElDWV9MT0FECiAgLSBBVURJVF9VU0VSX01BQ19DT05GSUdfQ0hBTkdFCi0KICBhY3Rpb246IGNoYW5nZWQtbG9naW4taWQtdG8KICBzdWJqZWN0OgogICAgcHJpbWFyeTogW29sZF9hdWlkLCBvbGQtYXVpZF0KICAgIHNlY29uZGFyeTogdWlkCiAgb2JqZWN0OgogICAgcHJpbWFyeTogYXVpZAogICAgd2hhdDogdXNlci1zZXNzaW9uCiAgcmVjb3JkX3R5cGVzOiBMT0dJTgotCiAgYWN0aW9uOiBtYWMtcGVybWlzc2lvbgogIHJlY29yZF90eXBlczogTUFDX0NIRUNLCi0KICBhY3Rpb246IGNoYW5nZWQtc2VsaW51eC1ib29sZWFuCiAgb2JqZWN0OgogICAgcHJpbWFyeTogYm9vbAogICAgd2hhdDogbWFjLWNvbmZpZwogIHJlY29yZF90eXBlczogTUFDX0NPTkZJR19DSEFOR0UKLQogIGFjdGlvbjogbG9hZGVkLXNlbGludXgtcG9saWN5CiAgb2JqZWN0OgogICAgd2hhdDogbWFjLWNvbmZpZwogIHJlY29yZF90eXBlczogTUFDX1BPTElDWV9MT0FECi0KICBhY3Rpb246IGNoYW5nZWQtc2VsaW51eC1lbmZvcmNlbWVudAogIG9iamVjdDoKICAgIHByaW1hcnk6IGVuZm9yY2luZwogICAgd2hhdDogbWFjLWNvbmZpZwogIHJlY29yZF90eXBlczogTUFDX1NUQVRVUwotCiAgYWN0aW9uOiBhc3NpZ25lZC11c2VyLXJvbGUtdG8KICBvYmplY3Q6CiAgICBwcmltYXJ5OiBbaWQsIGFjY3RdCiAgICB3aGF0OiBhY2NvdW50CiAgcmVjb3JkX3R5cGVzOiBST0xFX0FTU0lHTgotCiAgYWN0aW9uOiBtb2RpZmllZC1yb2xlCiAgcmVjb3JkX3R5cGVzOiBST0xFX01PRElGWQotCiAgYWN0aW9uOiByZW1vdmVkLXVzZS1yb2xlLWZyb20KICBvYmplY3Q6CiAgICBwcmltYXJ5OiBbaWQsIGFjY3RdCiAgICB3aGF0OiBhY2NvdW50CiAgcmVjb3JkX3R5cGVzOiBST0xFX1JFTU9WRQotCiAgYWN0aW9uOiB2aW9sYXRlZC1zZWNjb21wLXBvbGljeQogIG9iamVjdDoKICAgIHByaW1hcnk6IHN5c2NhbGwKICAgIHdoYXQ6IHByb2Nlc3MKICByZWNvcmRfdHlwZXM6IFNFQ0NPTVAKLQogIGFjdGlvbjogc3RhcnRlZC1zZXJ2aWNlCiAgb2JqZWN0OgogICAgcHJpbWFyeTogdW5pdAogICAgd2hhdDogc2VydmljZQogIHJlY29yZF90eXBlczogU0VSVklDRV9TVEFSVAotCiAgYWN0aW9uOiBzdG9wcGVkLXNlcnZpY2UKICBvYmplY3Q6CiAgICBwcmltYXJ5OiB1bml0CiAgICB3aGF0OiBzZXJ2aWNlCiAgcmVjb3JkX3R5cGVzOiBTRVJWSUNFX1NUT1AKLQogIGFjdGlvbjogYm9vdGVkLXN5c3RlbQogIG9iamVjdDoKICAgIHdoYXQ6IHN5c3RlbQogIHJlY29yZF90eXBlczogU1lTVEVNX0JPT1QKLQogIGFjdGlvbjogY2hhbmdlZC10by1ydW5sZXZlbAogIG9iamVjdDoKICAgIHByaW1hcnk6IG5ldy1sZXZlbAogICAgd2hhdDogc3lzdGVtCiAgcmVjb3JkX3R5cGVzOiBTWVNURU1fUlVOTEVWRUwKLQogIGFjdGlvbjogc2h1dGRvd24tc3lzdGVtCiAgb2JqZWN0OgogICAgd2hhdDogc3lzdGVtCiAgcmVjb3JkX3R5cGVzOiBTWVNURU1fU0hVVERPV04KLQogIGFjdGlvbjogc2VudC10ZXN0CiAgcmVjb3JkX3R5cGVzOiBURVNUCi0KICBhY3Rpb246IHVua25vd24KICByZWNvcmRfdHlwZXM6IFRSVVNURURfQVBQCi0KICBhY3Rpb246IHNlbnQtbWVzc2FnZQogIG9iamVjdDoKICAgIHByaW1hcnk6IGFkZHIKICByZWNvcmRfdHlwZXM6IFVTRVIKLQogIDw8OiAqbWFjcm8tdXNlci1zZXNzaW9uCiAgYWN0aW9uOiB3YXMtYXV0aG9yaXplZAogIHJlY29yZF90eXBlczogVVNFUl9BQ0NUCi0KICA8PDogKm1hY3JvLXVzZXItc2Vzc2lvbgogIGFjdGlvbjogYXV0aGVudGljYXRlZAogIHJlY29yZF90eXBlczogVVNFUl9BVVRICi0KICBhY3Rpb246IGFjY2Vzcy1wZXJtaXNzaW9uCiAgcmVjb3JkX3R5cGVzOiBVU0VSX0FWQwotCiAgPDw6ICptYWNyby11c2VyLXNlc3Npb24KICBhY3Rpb246IGNoYW5nZWQtcGFzc3dvcmQKICByZWNvcmRfdHlwZXM6IFVTRVJfQ0hBVVRIVE9LCi0KICBhY3Rpb246IHJhbi1jb21tYW5kCiAgb2JqZWN0OgogICAgcHJpbWFyeTogY21kCiAgICB3aGF0OiBwcm9jZXNzCiAgcmVjb3JkX3R5cGVzOiBVU0VSX0NNRAogIGRlc2NyaXB0aW9uOiA+CiAgICBUaGVzZSBtZXNzYWdlcyBhcmUgZnJvbSB1c2VyLXNwYWNlIGFwcHMsIGxpa2Ugc3VkbywgdGhhdCBsb2cgY29tbWFuZHMKICAgIGJlaW5nIHJ1biBieSBhIHVzZXIuIFRoZSB1aWQgY29udGFpbmVkIGluIHRoZXNlIG1lc3NhZ2VzIGlzIHVzZXIncyBVSUQgYXQKICAgIHRoZSB0aW1lIHRoZSBjb21tYW5kIHdhcyBydW4uIEl0IGlzIG5vdCB0aGUgInRhcmdldCIgVUlEIHVzZWQgdG8gcnVuIHRoZQogICAgY29tbWFuZCwgd2hpY2ggaXMgbm9ybWFsbHkgcm9vdC4KLQogIDw8OiAqbWFjcm8tdXNlci1zZXNzaW9uCiAgYWN0aW9uOiBlbmRlZC1zZXNzaW9uCiAgcmVjb3JkX3R5cGVzOiBVU0VSX0VORAotCiAgPDw6ICptYWNyby11c2VyLXNlc3Npb24KICBhY3Rpb246IGVycm9yCiAgcmVjb3JkX3R5cGVzOiBVU0VSX0VSUgogIHNvdXJjZV9pcDogW2FkZHJdCi0KICA8PDogKm1hY3JvLXVzZXItc2Vzc2lvbgogIGFjdGlvbjogbG9nZ2VkLWluCiAgcmVjb3JkX3R5cGVzOiBVU0VSX0xPR0lOCiAgc291cmNlX2lwOiBbYWRkcl0KLQogIDw8OiAqbWFjcm8tdXNlci1zZXNzaW9uCiAgYWN0aW9uOiBsb2dnZWQtb3V0CiAgcmVjb3JkX3R5cGVzOiBVU0VSX0xPR09VVAotCiAgYWN0aW9uOiBjaGFuZ2VkLW1hYy1jb25maWd1cmF0aW9uCiAgcmVjb3JkX3R5cGVzOiBVU0VSX01BQ19DT05GSUdfQ0hBTkdFCi0KICBhY3Rpb246IGxvYWRlZC1tYWMtcG9saWN5CiAgcmVjb3JkX3R5cGVzOiBVU0VSX01BQ19QT0xJQ1lfTE9BRAotCiAgPDw6ICptYWNyby11c2VyLXNlc3Npb24KICBhY3Rpb246IG1vZGlmaWVkLXVzZXItYWNjb3VudAogIHJlY29yZF90eXBlczogVVNFUl9NR01UCi0KICA8PDogKm1hY3JvLXVzZXItc2Vzc2lvbgogIGFjdGlvbjogY2hhbmdlZC1yb2xlLXRvCiAgb2JqZWN0OgogICAgcHJpbWFyeTogc2VsZWN0ZWQtY29udGV4dAogICAgd2hhdDogdXNlci1zZXNzaW9uCiAgcmVjb3JkX3R5cGVzOiBVU0VSX1JPTEVfQ0hBTkdFCi0KICBhY3Rpb246IGFjY2Vzcy1lcnJvcgogIHJlY29yZF90eXBlczogVVNFUl9TRUxJTlVYX0VSUgotCiAgPDw6ICptYWNyby11c2VyLXNlc3Npb24KICBhY3Rpb246IHN0YXJ0ZWQtc2Vzc2lvbgogIHJlY29yZF90eXBlczogVVNFUl9TVEFSVAogIHNvdXJjZV9pcDogW2FkZHJdCi0KICBhY3Rpb246IGNoYW5nZWQtY29uZmlndXJhdGlvbgogIG9iamVjdDoKICAgIHByaW1hcnk6IG9wCiAgICB3aGF0OiBzeXN0ZW0KICByZWNvcmRfdHlwZXM6IFVTWVNfQ09ORklHCi0KICBhY3Rpb246IGlzc3VlZC12bS1jb250cm9sCiAgb2JqZWN0OgogICAgcHJpbWFyeTogb3AKICAgIHNlY29uZGFyeTogdm0KICAgIHdoYXQ6IHZpcnR1YWwtbWFjaGluZQogIHJlY29yZF90eXBlczogVklSVF9DT05UUk9MCi0KICBhY3Rpb246IGNyZWF0ZWQtdm0taW1hZ2UKICByZWNvcmRfdHlwZXM6IFZJUlRfQ1JFQVRFCi0KICBhY3Rpb246IGRlbGV0ZWQtdm0taW1hZ2UKICByZWNvcmRfdHlwZXM6IFZJUlRfREVTVFJPWQotCiAgYWN0aW9uOiBjaGVja2VkLWludGVncml0eS1vZgogIHJlY29yZF90eXBlczogVklSVF9JTlRFR1JJVFlfQ0hFQ0sKLQogIGFjdGlvbjogYXNzaWduZWQtdm0taWQKICBvYmplY3Q6CiAgICBwcmltYXJ5OiB2bQogICAgd2hhdDogdmlydHVhbC1tYWNoaW5lCiAgcmVjb3JkX3R5cGVzOiBWSVJUX01BQ0hJTkVfSUQKLQogIGFjdGlvbjogbWlncmF0ZWQtdm0tZnJvbQogIHJlY29yZF90eXBlczogVklSVF9NSUdSQVRFX0lOCi0KICBhY3Rpb246IG1pZ3JhdGVkLXZtLXRvCiAgcmVjb3JkX3R5cGVzOiBWSVJUX01JR1JBVEVfT1VUCi0KICBhY3Rpb246IGFzc2lnbmVkLXZtLXJlc291cmNlCiAgb2JqZWN0OgogICAgcHJpbWFyeTogcmVzcmMKICAgIHNlY29uZGFyeTogdm0KICAgIHdoYXQ6IHZpcnR1YWwtbWFjaGluZQogIHJlY29yZF90eXBlczogVklSVF9SRVNPVVJDRQotIGFjdGlvbjogdHlwZWQKICBvYmplY3Q6CiAgICBwcmltYXJ5OiBkYXRhCiAgICB3aGF0OiBrZXlzdHJva2VzCiAgaG93OiBbY29tbSwgZXhlXQogIHJlY29yZF90eXBlczoKICAtIFRUWQogIC0gVVNFUl9UVFkK")
		assets["normalizationData"] = value
	}

	if value, found := assets[key]; found {
		return value, nil
	}
	return nil, fmt.Errorf("asset not found for key=%v", key)
}
