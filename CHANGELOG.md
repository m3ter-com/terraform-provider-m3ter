# Changelog

## 0.10.0 (2026-04-22)

Full Changelog: [v0.9.0...v0.10.0](https://github.com/m3ter-com/terraform-provider-m3ter/compare/v0.9.0...v0.10.0)

### Features

* add per-resource api permissions to schema description ([5cbb13e](https://github.com/m3ter-com/terraform-provider-m3ter/commit/5cbb13e0ef30767920a5063c54b042df244d2fee))


### Bug Fixes

* **ci:** in custom setup-go, pass through go-version and cache-dependency-path ([09aeb02](https://github.com/m3ter-com/terraform-provider-m3ter/commit/09aeb020e541c1ef3f86c4bfd1220fb4a716fa46))
* fall back to main branch if linking fails in CI ([d59623c](https://github.com/m3ter-com/terraform-provider-m3ter/commit/d59623cc95d03c86d6d2f12b21a4da5860c430a7))
* fix for failing to drop invalid module replace in link script ([1e99ca9](https://github.com/m3ter-com/terraform-provider-m3ter/commit/1e99ca9c1a8c075da033e0a215c15417642fde52))
* fix quoting typo ([ad24dff](https://github.com/m3ter-com/terraform-provider-m3ter/commit/ad24dff50c48eeee5550141784f680e3d41a2489))
* **tests:** update hc-install to fix PGP key mismatch ([40d1c9b](https://github.com/m3ter-com/terraform-provider-m3ter/commit/40d1c9b64c5456ea8f08edfbbc3c6650ddcc0c7c))


### Chores

* add local tmpfile directory ([219ecf0](https://github.com/m3ter-com/terraform-provider-m3ter/commit/219ecf0c9d833cb885d44cbff1fd26728090b1ff))
* **internal:** more robust bootstrap script ([6f23d72](https://github.com/m3ter-com/terraform-provider-m3ter/commit/6f23d724549b74a3fe851534c9158820275c3d2f))
* pin go releaser version ([5fbe44e](https://github.com/m3ter-com/terraform-provider-m3ter/commit/5fbe44e4f74c1c0b41cac7e9f3fc9b62a42733bd))
* **tests:** bump steady to v0.22.1 ([c8fedb2](https://github.com/m3ter-com/terraform-provider-m3ter/commit/c8fedb2ff35cfcade634b0e871c35198b12afcda))

## 0.9.0 (2026-03-31)

Full Changelog: [v0.8.0...v0.9.0](https://github.com/m3ter-com/terraform-provider-m3ter/compare/v0.8.0...v0.9.0)

### Features

* **api:** Updating OpenAPI Spec ([196fa85](https://github.com/m3ter-com/terraform-provider-m3ter/commit/196fa85cc8343af9bced43206e2feb31888dcf43))


### Bug Fixes

* improve linking behavior when developing on a branch not in the Go SDK ([091197e](https://github.com/m3ter-com/terraform-provider-m3ter/commit/091197e119016d4d0ba1a727e1383e8b8e708d86))
* improved workflow for developing on branches ([3d3bae6](https://github.com/m3ter-com/terraform-provider-m3ter/commit/3d3bae6c71b8773862825580b3c0e915dfabd4ae))
* no longer require an API key when building on production repos ([45f6bf6](https://github.com/m3ter-com/terraform-provider-m3ter/commit/45f6bf62b3fbdbd2c9db7bf795620e2655f2a0c1))
* patch style requests should never send empty json body for objects ([c070291](https://github.com/m3ter-com/terraform-provider-m3ter/commit/c070291fe0305442cc20d148d4d6fbd3546a0419))
* prefer named identifier field over id alias, missing ImportStates in certain resources ([8800b97](https://github.com/m3ter-com/terraform-provider-m3ter/commit/8800b9701ebc2c9aba8e79043b105344894ddcd1))
* spurious update plans for float attributes after import ([ef3024c](https://github.com/m3ter-com/terraform-provider-m3ter/commit/ef3024c37c9f93d2bc7a28b4981b988cf4d4a463))


### Chores

* **api:** minor updates ([c9f039c](https://github.com/m3ter-com/terraform-provider-m3ter/commit/c9f039c5af82f5be796d641017a7c4358f9e92ca))
* **docs:** add missing descriptions ([48b1dd4](https://github.com/m3ter-com/terraform-provider-m3ter/commit/48b1dd41dab557bb1023bcfd8427f0fd4b1673b0))
* **docs:** update terraform-plugin-docs to v0.24.0 ([530c3ef](https://github.com/m3ter-com/terraform-provider-m3ter/commit/530c3ef837e0c9bcd9d056406799c4300b3317b0))
* **internal:** codegen related update ([bdf9265](https://github.com/m3ter-com/terraform-provider-m3ter/commit/bdf926569fd0490520093854a1f69f12898fc3da))
* **internal:** codegen related update ([447b89e](https://github.com/m3ter-com/terraform-provider-m3ter/commit/447b89e57c7a862ec16df88ed84864cfcb4f73fd))
* **internal:** codegen related update ([976582b](https://github.com/m3ter-com/terraform-provider-m3ter/commit/976582baea9f2a44885b031f95fb4a73d6ab4ff1))
* **internal:** tweak CI branches ([d8b1441](https://github.com/m3ter-com/terraform-provider-m3ter/commit/d8b14413939d7db237c600a6167892e3e4e6a41c))
* **internal:** update gitignore ([4fea96f](https://github.com/m3ter-com/terraform-provider-m3ter/commit/4fea96f695e83aa7eb8a873e363788c4c2b9d422))
* **internal:** update multipart form array serialization ([32aa461](https://github.com/m3ter-com/terraform-provider-m3ter/commit/32aa4613d195e115d241eb8c236c88484106d08a))
* **test:** do not count install time for mock server timeout ([e5d0ee5](https://github.com/m3ter-com/terraform-provider-m3ter/commit/e5d0ee53a2c420271e91424bd83f1cb5f9af950b))
* **tests:** bump steady to v0.19.4 ([c686bc9](https://github.com/m3ter-com/terraform-provider-m3ter/commit/c686bc99b788edad08693939c0a069f5a0355534))
* **tests:** bump steady to v0.19.5 ([dc0be1c](https://github.com/m3ter-com/terraform-provider-m3ter/commit/dc0be1ce0ab4a26f73dcbcf92c56c4b75947a216))
* **tests:** bump steady to v0.19.6 ([30e0cf8](https://github.com/m3ter-com/terraform-provider-m3ter/commit/30e0cf80be038569d9ce08c09d2f222be2558a32))
* **tests:** bump steady to v0.19.7 ([0f258c4](https://github.com/m3ter-com/terraform-provider-m3ter/commit/0f258c440df83f2e02413b9309adf5ca11c854fc))
* **tests:** bump steady to v0.20.1 ([cd6fd1a](https://github.com/m3ter-com/terraform-provider-m3ter/commit/cd6fd1a10db45c47ece654638b54a7055ec1409e))
* **tests:** bump steady to v0.20.2 ([c0933ee](https://github.com/m3ter-com/terraform-provider-m3ter/commit/c0933eeee0b0470b6fba09c9bd5d38719f11a80c))


### Refactors

* **tests:** switch from prism to steady ([0c7c77a](https://github.com/m3ter-com/terraform-provider-m3ter/commit/0c7c77aa4745db36fcdaffad7c3cb1dff8581289))

## 0.8.0 (2026-01-29)

Full Changelog: [v0.7.0...v0.8.0](https://github.com/m3ter-com/terraform-provider-m3ter/compare/v0.7.0...v0.8.0)

### Features

* **api:** updating api spec + fixes ([f111743](https://github.com/m3ter-com/terraform-provider-m3ter/commit/f11174369fa310e7fde89576d391e51089e1f33f))


### Chores

* bump dependency version ([d0c5dfd](https://github.com/m3ter-com/terraform-provider-m3ter/commit/d0c5dfd0a3e114a117d98decf70ff3e7b7a2ae32))
* **internal:** codegen related update ([c386327](https://github.com/m3ter-com/terraform-provider-m3ter/commit/c386327bc7f768a0f6afea010258beebe3a5edae))
* **internal:** codegen related update ([2c702e8](https://github.com/m3ter-com/terraform-provider-m3ter/commit/2c702e8c2f97fa3c7e266963a48474a2752d659a))
* **internal:** codegen related update ([5bdb5fd](https://github.com/m3ter-com/terraform-provider-m3ter/commit/5bdb5fdb923996e9a86ebece5e8efe4e9bc6b335))
* **internal:** update `actions/checkout` version ([9698735](https://github.com/m3ter-com/terraform-provider-m3ter/commit/9698735be1c5808d233a75330664c7ebfe112ed4))
* update Go SDK version ([f6d8d10](https://github.com/m3ter-com/terraform-provider-m3ter/commit/f6d8d10afb2cdc013691ac1660b947b515c156b4))

## 0.7.0 (2025-12-03)

Full Changelog: [v0.6.1...v0.7.0](https://github.com/m3ter-com/terraform-provider-m3ter/compare/v0.6.1...v0.7.0)

### Features

* **api:** spec update ([7ddb4c1](https://github.com/m3ter-com/terraform-provider-m3ter/commit/7ddb4c1aa22e96d19a3469779004ddebf7eb35b1))


### Bug Fixes

* **client:** correctly encode map patches ([01b0927](https://github.com/m3ter-com/terraform-provider-m3ter/commit/01b0927ae5c35aef1cfe6d7d2205bb3e5ccfe7fd))
* **client:** correctly patch `null` -&gt; zero value ([fc9ac4a](https://github.com/m3ter-com/terraform-provider-m3ter/commit/fc9ac4a357f5b9cf934c419f55d251da8a925212))
* data sources should handle optional parameters specified by the provider ([bd60bb7](https://github.com/m3ter-com/terraform-provider-m3ter/commit/bd60bb7fb3cf6310639dc3b473ded66622df090f))
* ensure dynamic values always yield valid container inner values ([e69b26a](https://github.com/m3ter-com/terraform-provider-m3ter/commit/e69b26a36b53a555e4ae639770ccd328567b914d))
* list style data sources should always have id value populated ([904c404](https://github.com/m3ter-com/terraform-provider-m3ter/commit/904c404f5474c21fda2c946d541f40dc4014f05e))


### Chores

* ensure tests build as part of lint step ([9b67030](https://github.com/m3ter-com/terraform-provider-m3ter/commit/9b67030836e1a53fb821f7a1d87b6eaf391d784b))
* **internal:** address linter warnings ([98388fe](https://github.com/m3ter-com/terraform-provider-m3ter/commit/98388fe7f9f836e6aef33e001b75ce3d5fd40810))
* **internal:** codegen related update ([2efb329](https://github.com/m3ter-com/terraform-provider-m3ter/commit/2efb329b0d2adad97af0a81ec1a095b13f2999f2))
* **internal:** refactor the apijson encoder ([ede372b](https://github.com/m3ter-com/terraform-provider-m3ter/commit/ede372be65e79475dc6ef58557c39f783001cb4b))
* **internal:** update `interface{}` to `any` ([ac25487](https://github.com/m3ter-com/terraform-provider-m3ter/commit/ac25487aa2ab1cd0bc6835cbbd4e35fa61ad74f2))

## 0.6.1 (2025-10-08)

Full Changelog: [v0.6.0...v0.6.1](https://github.com/m3ter-com/terraform-provider-m3ter/compare/v0.6.0...v0.6.1)

### Features

* **docs:** Update examples for singleton resources ([3c5540a](https://github.com/m3ter-com/terraform-provider-m3ter/commit/3c5540a7c96ba56b59a0f621b306d50386621142))

## 0.6.0 (2025-10-07)

Full Changelog: [v0.5.0...v0.6.0](https://github.com/m3ter-com/terraform-provider-m3ter/compare/v0.5.0...v0.6.0)

### Features

* added capability for `dynamicvalidator` to do arbitrary semantic equivalence check ([ecc974d](https://github.com/m3ter-com/terraform-provider-m3ter/commit/ecc974dff52ac8a8860c33331522d055ea6a5644))
* **internal:** support CustomMarshaler interface for encoding types ([a853a5e](https://github.com/m3ter-com/terraform-provider-m3ter/commit/a853a5e92e55adfd84060bd8fdf257a77269a342))
* **sdk:** Support singleton resources in Terraform ([1aeb866](https://github.com/m3ter-com/terraform-provider-m3ter/commit/1aeb8665eac44a151f4a985cfe55355b1e5af218))


### Bug Fixes

* bugfix for setting JSON keys with special characters ([6cdc552](https://github.com/m3ter-com/terraform-provider-m3ter/commit/6cdc552966d54e95a6b831563991296565d584c2))
* correctly detect more ID attributes for data sources ([44fbd84](https://github.com/m3ter-com/terraform-provider-m3ter/commit/44fbd84d43b69e01cb01afb778daa28e10b1e109))
* read by id data sources should have required IDs ([8a09d26](https://github.com/m3ter-com/terraform-provider-m3ter/commit/8a09d26cbc8d6de0b5fef0262d2f16d87617099b))


### Chores

* do not install brew dependencies in ./scripts/bootstrap by default ([bbd3397](https://github.com/m3ter-com/terraform-provider-m3ter/commit/bbd33972f7406008feb30f975200e53f8abd4210))
* ensure `tfplugindocs` always use `/var/tmp` for compilation on linux ([4855dcc](https://github.com/m3ter-com/terraform-provider-m3ter/commit/4855dcc9a89d41ba7336f28ec760e9c624a96f31))

## 0.5.0 (2025-08-29)

Full Changelog: [v0.4.0...v0.5.0](https://github.com/m3ter-com/terraform-provider-m3ter/compare/v0.4.0...v0.5.0)

### Features

* **api:** Add examples to integrations and events ([a8addd1](https://github.com/m3ter-com/terraform-provider-m3ter/commit/a8addd19aed771b7828f7c5c06e1cd79a6b2aa75))
* **api:** Bump Go SDK in Terraform ([3a21293](https://github.com/m3ter-com/terraform-provider-m3ter/commit/3a212935683a9b9e2ed14652b6a3147616249518))
* ensure `internal/apiform` encoder can handle "force_encode" serialization tag ([ba959f0](https://github.com/m3ter-com/terraform-provider-m3ter/commit/ba959f025041884fb1d80171571a03fae122ae19))


### Bug Fixes

* **api:** handle mismatched dynamic array types in state and plan during serialization ([da59eff](https://github.com/m3ter-com/terraform-provider-m3ter/commit/da59eff5ad857f0f48f98718093670aaf87814d5))
* check before de-referencing potentially null ptrs ([a0216da](https://github.com/m3ter-com/terraform-provider-m3ter/commit/a0216da8b0e0b36cdad2ec40e2331af1448882ed))
* dynamic type validators should handle int and floats correctly ([c29d0ff](https://github.com/m3ter-com/terraform-provider-m3ter/commit/c29d0ff1172defda4443517800452cf64b2ff88a))
* encoder crash for nested nils in dynamic types ([2d4c2d0](https://github.com/m3ter-com/terraform-provider-m3ter/commit/2d4c2d001842a110ae8bb4e52c361f55e0e42d64))
* populate computed_optional collections from API responses ([af5d013](https://github.com/m3ter-com/terraform-provider-m3ter/commit/af5d013ebbdd52ecbc2fbd345dc46859d3c38686))
* properly handle null nested objects in customfield marshaling ([e3c74dd](https://github.com/m3ter-com/terraform-provider-m3ter/commit/e3c74dd15d3eabd4609db55e5e228fd67dededa3))


### Chores

* improve integrity test error messages ([776948d](https://github.com/m3ter-com/terraform-provider-m3ter/commit/776948d3f094c422231e522b38ea29160b2e4075))
* **internal:** add test rule to lint for dynamic attributes that do not have planmodifier ([475d30e](https://github.com/m3ter-com/terraform-provider-m3ter/commit/475d30e24face3e0ef16a3574b8e2ad3dec817d0))
* **internal:** codegen related update ([bf68f01](https://github.com/m3ter-com/terraform-provider-m3ter/commit/bf68f01f3d8b4c8eaf8e966f5428e3efeadff9fc))
* **internal:** upgrade cloudflare/circl ([e13b6c1](https://github.com/m3ter-com/terraform-provider-m3ter/commit/e13b6c18ed41efbd8086908b8f9eb25f95c3ddba))
* update @stainless-api/prism-cli to v5.15.0 ([d8a5781](https://github.com/m3ter-com/terraform-provider-m3ter/commit/d8a5781869f725ceb109c319dedf13b688399791))

## 0.4.0 (2025-07-28)

Full Changelog: [v0.3.2...v0.4.0](https://github.com/m3ter-com/terraform-provider-m3ter/compare/v0.3.2...v0.4.0)

### Features

* **api:** add aggregationId query param to ListPricings ([f69c5d5](https://github.com/m3ter-com/terraform-provider-m3ter/commit/f69c5d59a0e4b1334c9e6ad889eba55d090b1792))
* **api:** remove audit fields from Terraform provider ([2aea686](https://github.com/m3ter-com/terraform-provider-m3ter/commit/2aea6867cf9b91bbfac0295673763b13df77c173))


### Chores

* **internal:** codegen related update ([1cd15bb](https://github.com/m3ter-com/terraform-provider-m3ter/commit/1cd15bb9ba826dbe59c415f26eef93d336a9be05))

## 0.3.2 (2025-07-18)

Full Changelog: [v0.3.1...v0.3.2](https://github.com/m3ter-com/terraform-provider-m3ter/compare/v0.3.1...v0.3.2)

### Chores

* **internal:** version bump ([d92be99](https://github.com/m3ter-com/terraform-provider-m3ter/commit/d92be996ebe1019d41457d1c730c101f43c40290))

## 0.3.1 (2025-07-18)

Full Changelog: [v0.3.0...v0.3.1](https://github.com/m3ter-com/terraform-provider-m3ter/compare/v0.3.0...v0.3.1)

### Bug Fixes

* **api:** remove custom fields from Terraform ([20d4324](https://github.com/m3ter-com/terraform-provider-m3ter/commit/20d43240c85c8510dca20b3cb68964e38d7bc098))
* **api:** use WebhookDestinationResponse in webhook endpoint response schemas ([432a258](https://github.com/m3ter-com/terraform-provider-m3ter/commit/432a2585488c2d84629b9ff542f96dcbe25e6734))

## 0.3.0 (2025-07-10)

Full Changelog: [v0.2.0-alpha...v0.3.0](https://github.com/m3ter-com/terraform-provider-m3ter/compare/v0.2.0-alpha...v0.3.0)

### Features

* **api:** add x-stainless-terraform-always-send=true to version attributes ([3fbaf3a](https://github.com/m3ter-com/terraform-provider-m3ter/commit/3fbaf3a8ea4febca930461c01ef0ac2a4cc6535c))
* new option to send computed values back to server ([12950a0](https://github.com/m3ter-com/terraform-provider-m3ter/commit/12950a042494c3102723ac4df47f4ae14d04b846))

## 0.2.0-alpha (2025-07-09)

Full Changelog: [v0.1.1-alpha...v0.2.0-alpha](https://github.com/m3ter-com/terraform-provider-m3ter/compare/v0.1.1-alpha...v0.2.0-alpha)

### Features

* **api:** Bump Go SDK version in Terraform provider ([a6876cc](https://github.com/m3ter-com/terraform-provider-m3ter/commit/a6876cc42efd4974753a2f2201c0ed77f2ea9169))

## 0.1.1-alpha (2025-07-09)

Full Changelog: [v0.1.0-alpha...v0.1.1-alpha](https://github.com/m3ter-com/terraform-provider-m3ter/compare/v0.1.0-alpha...v0.1.1-alpha)

## 0.1.0-alpha (2025-07-09)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha](https://github.com/m3ter-com/terraform-provider-m3ter/compare/v0.1.0-alpha.1...v0.1.0-alpha)

### Features

* **api:** mark `version` attribute as computed in Terraform ([ed6889c](https://github.com/m3ter-com/terraform-provider-m3ter/commit/ed6889cf6d774c8308a644db52632dc4c16d23af))
* **api:** Set "version" attribute to "computed_optional" in Terraform ([ab7aeff](https://github.com/m3ter-com/terraform-provider-m3ter/commit/ab7aeff6a2c0b507b3cc492001894e0e7c80ce1a))


### Bug Fixes

* **ci:** release-doctor — report correct token name ([f483287](https://github.com/m3ter-com/terraform-provider-m3ter/commit/f4832877b2ac41187aeb25025f62f7cffc58915a))
* fix attribute definitions for collections of dynamic values ([c14684a](https://github.com/m3ter-com/terraform-provider-m3ter/commit/c14684a911eec8e723f0711aa8deaf649a92816c))
* null nested attribute decoding ([55aedf4](https://github.com/m3ter-com/terraform-provider-m3ter/commit/55aedf4e6c93ee1973de7ce0ee1fd82af200b7e1))


### Chores

* **ci:** only run for pushes and fork pull requests ([6b210fe](https://github.com/m3ter-com/terraform-provider-m3ter/commit/6b210fee51cd2248e94f59ee12cd5d08079023de))
* **internal:** codegen related update ([3245fe0](https://github.com/m3ter-com/terraform-provider-m3ter/commit/3245fe0387cc50dd36ab3c9d36aa6fdf0b59b907))
* **internal:** codegen related update ([017ecf1](https://github.com/m3ter-com/terraform-provider-m3ter/commit/017ecf1992be4b33d27c881d5517b87df8b9f44e))
* **internal:** codegen related update ([20c2d23](https://github.com/m3ter-com/terraform-provider-m3ter/commit/20c2d23fd6f9994323a1b3450df6c98a4d748a14))
* **internal:** update examples ([edbd61c](https://github.com/m3ter-com/terraform-provider-m3ter/commit/edbd61c381207c06822ed23c6160613e70bfd5b2))

## 0.1.0-alpha.1 (2025-06-17)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/m3ter-com/terraform-provider-m3ter/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### ⚠ BREAKING CHANGES

* breaking change - singluar data sources should not have squashed schemas ([#2](https://github.com/m3ter-com/terraform-provider-m3ter/issues/2))

### Features

* add doc string to specify what legal terraform values are for enums ([#31](https://github.com/m3ter-com/terraform-provider-m3ter/issues/31)) ([8dc2f4f](https://github.com/m3ter-com/terraform-provider-m3ter/commit/8dc2f4f8a2812e6b0b392b4bbd8cbbe78a51a14f))
* add docs generation to format script ([#41](https://github.com/m3ter-com/terraform-provider-m3ter/issues/41)) ([ed78af0](https://github.com/m3ter-com/terraform-provider-m3ter/commit/ed78af077db3e5530e035d430d2d04fd37646532))
* add SKIP_BREW env var to ./scripts/bootstrap ([#40](https://github.com/m3ter-com/terraform-provider-m3ter/issues/40)) ([38c6026](https://github.com/m3ter-com/terraform-provider-m3ter/commit/38c6026dd01db5cc83490f2fb1eb2c120baa31a6))
* **api:** add measurement request model ([a42e4a6](https://github.com/m3ter-com/terraform-provider-m3ter/commit/a42e4a6038c52f8053681f4ccc009f69e0663e2d))
* **api:** add missing endpoints ([#29](https://github.com/m3ter-com/terraform-provider-m3ter/issues/29)) ([b629660](https://github.com/m3ter-com/terraform-provider-m3ter/commit/b62966084edf4ecd22891f0cc29924bf3df0911e))
* **api:** add more endpoints ([#18](https://github.com/m3ter-com/terraform-provider-m3ter/issues/18)) ([fbfb374](https://github.com/m3ter-com/terraform-provider-m3ter/commit/fbfb3746ae90393153dabb6e386fe227c2e0ed31))
* **api:** add orgId path param to client settings ([#27](https://github.com/m3ter-com/terraform-provider-m3ter/issues/27)) ([51efbcb](https://github.com/m3ter-com/terraform-provider-m3ter/commit/51efbcb9a01deedb362368bcd3a826e90b580e5f))
* **api:** add statements ([7a366a1](https://github.com/m3ter-com/terraform-provider-m3ter/commit/7a366a188e894471adec96b65cd71cb513c011b4))
* **api:** Config update ([#12](https://github.com/m3ter-com/terraform-provider-m3ter/issues/12)) ([e1115df](https://github.com/m3ter-com/terraform-provider-m3ter/commit/e1115df2cab39074b9d1cf65401af51db073e314))
* **api:** create ad hoc data export endpoint ([#24](https://github.com/m3ter-com/terraform-provider-m3ter/issues/24)) ([3c93af5](https://github.com/m3ter-com/terraform-provider-m3ter/commit/3c93af5cc25f6765ebf54d7773a4e0f1008e7691))
* **api:** display deprecation messages on resources and attributes ([#47](https://github.com/m3ter-com/terraform-provider-m3ter/issues/47)) ([ae2b0b4](https://github.com/m3ter-com/terraform-provider-m3ter/commit/ae2b0b474bd500bec9e4b7e1f9dfbb21c579344c))
* **api:** fixing warnings ([#38](https://github.com/m3ter-com/terraform-provider-m3ter/issues/38)) ([9512232](https://github.com/m3ter-com/terraform-provider-m3ter/commit/9512232e38eafbb269cbb074be9e4d0ee4b09fa3))
* **api:** Introduce OrganizationConfigRequest model ([389d94b](https://github.com/m3ter-com/terraform-provider-m3ter/commit/389d94b1db34c32a05e261b2f2b924aca8422036))
* **api:** manual updates ([#39](https://github.com/m3ter-com/terraform-provider-m3ter/issues/39)) ([acf74a5](https://github.com/m3ter-com/terraform-provider-m3ter/commit/acf74a5b7d109e6018bf7a58cf79d2f2cb4107cc))
* **api:** OpenAPI spec update ([fca0b3b](https://github.com/m3ter-com/terraform-provider-m3ter/commit/fca0b3b9bb112b74f6a5e6f473907593ec7e9f7a))
* **api:** snake case method names ([#30](https://github.com/m3ter-com/terraform-provider-m3ter/issues/30)) ([dfac597](https://github.com/m3ter-com/terraform-provider-m3ter/commit/dfac5972fc829dd11b3fa57b8970b40c1e92c3da))
* **api:** Spec fixes ([10cabf8](https://github.com/m3ter-com/terraform-provider-m3ter/commit/10cabf8047a24c715ac1905c326081204d44bc5a))
* **api:** Spec Update + Various Fixes ([#35](https://github.com/m3ter-com/terraform-provider-m3ter/issues/35)) ([4f6780a](https://github.com/m3ter-com/terraform-provider-m3ter/commit/4f6780a0e50ec61f349730fbf45289381d56caf0))
* **api:** support shared path parameters being set on the provider ([#46](https://github.com/m3ter-com/terraform-provider-m3ter/issues/46)) ([2dabc4a](https://github.com/m3ter-com/terraform-provider-m3ter/commit/2dabc4a14705f1ade747047d60f70189fced5739))
* **api:** update contact email and package name ([#49](https://github.com/m3ter-com/terraform-provider-m3ter/issues/49)) ([44b9bdd](https://github.com/m3ter-com/terraform-provider-m3ter/commit/44b9bdd69139692872b17687c84ea8bb88bf8a7a))
* **api:** Update custom field type information ([#15](https://github.com/m3ter-com/terraform-provider-m3ter/issues/15)) ([60fc22e](https://github.com/m3ter-com/terraform-provider-m3ter/commit/60fc22ef237211228a1f89b854533457b6fdb4b9))
* **api:** update open api spec ([df131eb](https://github.com/m3ter-com/terraform-provider-m3ter/commit/df131ebdbc71be3895a7196bdd1861a9ea8c7742))
* **api:** update open api spec ([#28](https://github.com/m3ter-com/terraform-provider-m3ter/issues/28)) ([48ece47](https://github.com/m3ter-com/terraform-provider-m3ter/commit/48ece47f1e80210498ac97d505fbdf65da910e61))
* **api:** update OpenAPI spec + associated fixes ([a3e64db](https://github.com/m3ter-com/terraform-provider-m3ter/commit/a3e64db12b7d9f8ce1586ff49b09b1f08e19bdba))
* **api:** updated OpenAPI spec ([#10](https://github.com/m3ter-com/terraform-provider-m3ter/issues/10)) ([d581d23](https://github.com/m3ter-com/terraform-provider-m3ter/commit/d581d23b2c6ccf06ad0363ba3d294f6e2bcf6d03))
* breaking change - singluar data sources should not have squashed schemas ([#2](https://github.com/m3ter-com/terraform-provider-m3ter/issues/2)) ([4272741](https://github.com/m3ter-com/terraform-provider-m3ter/commit/4272741f2e2d441cd0893dce038adf0c29045627))
* **build:** allow for building against private go repos ([#17](https://github.com/m3ter-com/terraform-provider-m3ter/issues/17)) ([49df86a](https://github.com/m3ter-com/terraform-provider-m3ter/commit/49df86aec8499d85c45894e7d400203485dec959))
* **client:** support environments property from Stainless config ([d6fa44b](https://github.com/m3ter-com/terraform-provider-m3ter/commit/d6fa44bbf2a23208ed446347bd15a483b735defd))
* **docs:** improve readme ([#9](https://github.com/m3ter-com/terraform-provider-m3ter/issues/9)) ([617aa10](https://github.com/m3ter-com/terraform-provider-m3ter/commit/617aa105de011fa43d8cc8ad6d479c0fc8c32254))


### Bug Fixes

* **api:** better support for environment variables as provider properties ([#43](https://github.com/m3ter-com/terraform-provider-m3ter/issues/43)) ([c3ea5fa](https://github.com/m3ter-com/terraform-provider-m3ter/commit/c3ea5facb4e73157135b9f846cab35993e1dbe51))
* **api:** improve drift detection for resources ([#13](https://github.com/m3ter-com/terraform-provider-m3ter/issues/13)) ([748b63c](https://github.com/m3ter-com/terraform-provider-m3ter/commit/748b63cdee8fec9ff98c1a4f50ad5c4f2764b976))
* **api:** improve drift detection for resources ([#26](https://github.com/m3ter-com/terraform-provider-m3ter/issues/26)) ([012542f](https://github.com/m3ter-com/terraform-provider-m3ter/commit/012542f23d2c0f1e6f6a72df026948c5be190a6f))
* **api:** skip generation of update endpoint when only a create endpoint exists ([#45](https://github.com/m3ter-com/terraform-provider-m3ter/issues/45)) ([0f4696c](https://github.com/m3ter-com/terraform-provider-m3ter/commit/0f4696c7508921601946212161c9e9b729d98d51))
* **api:** terraform release readiness ([fd5efe9](https://github.com/m3ter-com/terraform-provider-m3ter/commit/fd5efe9a8d42e0177af8e8008b53d9d0e4d73857))
* **build:** do not fail if go mod tidy fails during bootstrapping ([c506f24](https://github.com/m3ter-com/terraform-provider-m3ter/commit/c506f246692c29b78bcaebbd031e9203a1e544fb))
* **build:** enable building against private Go production repos ([fcc4eb9](https://github.com/m3ter-com/terraform-provider-m3ter/commit/fcc4eb92c46093db529ff9eab2c8c5310fd52e55))
* **build:** ensure scripts/generate-docs works regardless of PATH ([#6](https://github.com/m3ter-com/terraform-provider-m3ter/issues/6)) ([7531196](https://github.com/m3ter-com/terraform-provider-m3ter/commit/7531196cf60208b330e32012587dc0789f020009))
* **build:** improve release process ([#4](https://github.com/m3ter-com/terraform-provider-m3ter/issues/4)) ([32761fb](https://github.com/m3ter-com/terraform-provider-m3ter/commit/32761fb10add520144b2ff0e00273cb2e2ccf7fd))
* **codegen:** prevent spurious resource replacement due to computed_optional properties ([4863e43](https://github.com/m3ter-com/terraform-provider-m3ter/commit/4863e43e81014a47bd3de79ce553ee909c109486))
* data sources' `find_one_by` functionality should re-use plural data sources' code path ([#11](https://github.com/m3ter-com/terraform-provider-m3ter/issues/11)) ([c1f40bf](https://github.com/m3ter-com/terraform-provider-m3ter/commit/c1f40bf082ea402cd4f1646141bed4ff78d047ae))
* **datasource:** improve configurability of path parameters on data sources ([#25](https://github.com/m3ter-com/terraform-provider-m3ter/issues/25)) ([e4bbaa8](https://github.com/m3ter-com/terraform-provider-m3ter/commit/e4bbaa8c41d0daff6a1237ab009ca7a49d62a756))
* deduplicate unknown entries in union ([#8](https://github.com/m3ter-com/terraform-provider-m3ter/issues/8)) ([caf50dd](https://github.com/m3ter-com/terraform-provider-m3ter/commit/caf50dd750a494891f5810830e4af024b8423327))
* do not call path.Base on ContentType ([#21](https://github.com/m3ter-com/terraform-provider-m3ter/issues/21)) ([8f145a8](https://github.com/m3ter-com/terraform-provider-m3ter/commit/8f145a86a6c6eb15bee2ad8db0451b830b983a4b))
* **docs:** ensure schema docstrings always match the correct schema ([65df914](https://github.com/m3ter-com/terraform-provider-m3ter/commit/65df914bcbd804830e8df4690f2cda8c8dc002b8))
* fix caching issue between Unmarshal and UnmarshalComputed ([dab6698](https://github.com/m3ter-com/terraform-provider-m3ter/commit/dab66982afe7f0eb2b4e1fbe5e6fa836ca944041))
* **internal:** more consistent handling of terraform attribute names ([1cd2411](https://github.com/m3ter-com/terraform-provider-m3ter/commit/1cd2411210a8227b53a6b1e64d8a201f51a8c0a6))
* only unmarshal attributes that exist on the read response schema during refresh ([1ee6255](https://github.com/m3ter-com/terraform-provider-m3ter/commit/1ee6255f75806e4d5e4d396f9443ae5580b4a903))
* **release:** update README and version correctly in release PRs ([da0e2fe](https://github.com/m3ter-com/terraform-provider-m3ter/commit/da0e2fe55676294a483a07f8d1b40dd232adf9af))
* **schema:** fix configurability calculation for nested properties without computed values ([3fae90d](https://github.com/m3ter-com/terraform-provider-m3ter/commit/3fae90d3d3f407d76bd3b14e4e1a0c51371900a0))


### Chores

* **build:** scripts/format should not fail if generate-docs fails ([#48](https://github.com/m3ter-com/terraform-provider-m3ter/issues/48)) ([78bd41f](https://github.com/m3ter-com/terraform-provider-m3ter/commit/78bd41fc4210d281f3ca5bcb1468cd8b73bfa1ff))
* **build:** update go.mod indirect dependencies ([6273d9d](https://github.com/m3ter-com/terraform-provider-m3ter/commit/6273d9dfae27000d83952f4bf3853f600a17dd4a))
* bump deps to avoid GetResourceIdentitySchemas errors for Terraform CLI v1.12+ ([416ecbf](https://github.com/m3ter-com/terraform-provider-m3ter/commit/416ecbf750718e109adcc4a8d2cb6eefcab6f2e5))
* casing change in doc string ([#32](https://github.com/m3ter-com/terraform-provider-m3ter/issues/32)) ([528e541](https://github.com/m3ter-com/terraform-provider-m3ter/commit/528e541e70b503c73ab620b9d31669575046eb0f))
* **ci:** enable for pull requests ([018ffd1](https://github.com/m3ter-com/terraform-provider-m3ter/commit/018ffd1241495afacfcd1d13a36c65ec75544dac))
* **ci:** only use depot for staging repos ([901c9ed](https://github.com/m3ter-com/terraform-provider-m3ter/commit/901c9edf540916d11e142abb7ede68deeaa6e7f2))
* **ci:** run on more branches ([0f73feb](https://github.com/m3ter-com/terraform-provider-m3ter/commit/0f73febc700092e41b2b81355bb9069b705ea593))
* **ci:** run on more branches and use depot runners ([1ab8bc0](https://github.com/m3ter-com/terraform-provider-m3ter/commit/1ab8bc06b44dd739ebe666f4f364fa7287d49768))
* **docs:** grammar improvements ([a64969e](https://github.com/m3ter-com/terraform-provider-m3ter/commit/a64969e7b8431ae8e4c5fc2a869a894b4830be1e))
* go live ([e0b9132](https://github.com/m3ter-com/terraform-provider-m3ter/commit/e0b9132f953eb330bea03e267f93ac6e0f410eeb))
* **internal:** codegen related update ([4fc273a](https://github.com/m3ter-com/terraform-provider-m3ter/commit/4fc273ab408a1f93dac4aa3f55ddf405b6fc3580))
* **internal:** codegen related update ([03eed97](https://github.com/m3ter-com/terraform-provider-m3ter/commit/03eed97bfad35c415ced03c109cd7a0ddb13ed81))
* **internal:** codegen related update ([3843cd3](https://github.com/m3ter-com/terraform-provider-m3ter/commit/3843cd3058bc20f261148aaf975f7a76c2533acc))
* **internal:** codegen related update ([222ddf8](https://github.com/m3ter-com/terraform-provider-m3ter/commit/222ddf89a2490f0f61cb8389c5de8c08942ef3a7))
* **internal:** codegen related update ([#14](https://github.com/m3ter-com/terraform-provider-m3ter/issues/14)) ([83145f2](https://github.com/m3ter-com/terraform-provider-m3ter/commit/83145f2f94fa8dbab842e9fb3083b4dcba30f595))
* **internal:** codegen related update ([#19](https://github.com/m3ter-com/terraform-provider-m3ter/issues/19)) ([b52af8d](https://github.com/m3ter-com/terraform-provider-m3ter/commit/b52af8d9afd8f38838c13e9f196ffd9229efad0a))
* **internal:** codegen related update ([#22](https://github.com/m3ter-com/terraform-provider-m3ter/issues/22)) ([3efb500](https://github.com/m3ter-com/terraform-provider-m3ter/commit/3efb5002586d48b70d7780cd4440c49b69bd1b53))
* **internal:** codegen related update ([#23](https://github.com/m3ter-com/terraform-provider-m3ter/issues/23)) ([9706f88](https://github.com/m3ter-com/terraform-provider-m3ter/commit/9706f88863ee2632e76e8057ed555e31e7277be7))
* **internal:** codegen related update ([#3](https://github.com/m3ter-com/terraform-provider-m3ter/issues/3)) ([1f0febb](https://github.com/m3ter-com/terraform-provider-m3ter/commit/1f0febb8abbe93dfef8354a06bd25d5abe4660cc))
* **internal:** codegen related update ([#42](https://github.com/m3ter-com/terraform-provider-m3ter/issues/42)) ([3cb51ff](https://github.com/m3ter-com/terraform-provider-m3ter/commit/3cb51ff3e5f76f289f278d347651910fd19d849e))
* **internal:** codegen related update ([#5](https://github.com/m3ter-com/terraform-provider-m3ter/issues/5)) ([f28474f](https://github.com/m3ter-com/terraform-provider-m3ter/commit/f28474f643caa05aa5930289ea96631550a4ecde))
* **internal:** codegen related update ([#51](https://github.com/m3ter-com/terraform-provider-m3ter/issues/51)) ([a3c4cbd](https://github.com/m3ter-com/terraform-provider-m3ter/commit/a3c4cbd875fabf1e20011019b9bad698da749d95))
* **internal:** codegen related update ([#7](https://github.com/m3ter-com/terraform-provider-m3ter/issues/7)) ([65b4ad0](https://github.com/m3ter-com/terraform-provider-m3ter/commit/65b4ad0cff3885ca5fd914a1bcb5089eb6c5a829))
* **internal:** fix tests ([67cfd40](https://github.com/m3ter-com/terraform-provider-m3ter/commit/67cfd40d454263e0facb607010e60fe51e0370fe))
* **internal:** updates ([cd513e6](https://github.com/m3ter-com/terraform-provider-m3ter/commit/cd513e6d29e86dff06ae8c3efe9d213973036cd6))
* minor change to tests ([#20](https://github.com/m3ter-com/terraform-provider-m3ter/issues/20)) ([4425d2f](https://github.com/m3ter-com/terraform-provider-m3ter/commit/4425d2f0ac73284f53d26a7fa6858521c4f15bcb))
* org ID at the client level is required ([#37](https://github.com/m3ter-com/terraform-provider-m3ter/issues/37)) ([e4ddbf1](https://github.com/m3ter-com/terraform-provider-m3ter/commit/e4ddbf1bcd9218d6dc4d6c718456c495694618c3))
* org ID client arg is optional ([#36](https://github.com/m3ter-com/terraform-provider-m3ter/issues/36)) ([3dd283b](https://github.com/m3ter-com/terraform-provider-m3ter/commit/3dd283b660956a092cbcce99c897c36430aca018))
* **release:** enable release PRs to be opened before publishing to hashicorp ([5433921](https://github.com/m3ter-com/terraform-provider-m3ter/commit/543392118d68ac8dcd2e19d005bafe60738c21a7))
* simplify string literals ([#34](https://github.com/m3ter-com/terraform-provider-m3ter/issues/34)) ([9a90193](https://github.com/m3ter-com/terraform-provider-m3ter/commit/9a90193c2edc14ed23d547dbd83b585c0f1e8099))
* update dependencies ([#44](https://github.com/m3ter-com/terraform-provider-m3ter/issues/44)) ([afafa30](https://github.com/m3ter-com/terraform-provider-m3ter/commit/afafa3044b5bf7639adc843f14850d0ffdca2068))
* update SDK settings ([#1](https://github.com/m3ter-com/terraform-provider-m3ter/issues/1)) ([7b9f150](https://github.com/m3ter-com/terraform-provider-m3ter/commit/7b9f1509853abd8253cded70c039bff067912232))


### Documentation

* update URLs from stainlessapi.com to stainless.com ([#33](https://github.com/m3ter-com/terraform-provider-m3ter/issues/33)) ([b7a0140](https://github.com/m3ter-com/terraform-provider-m3ter/commit/b7a0140b350226d71b271358885f5dff386a33a9))
* Use "My Org Id" in example requests ([#50](https://github.com/m3ter-com/terraform-provider-m3ter/issues/50)) ([3d581be](https://github.com/m3ter-com/terraform-provider-m3ter/commit/3d581bed453688d08f4f206738fdce79c2006335))
