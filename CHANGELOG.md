# Changelog

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
