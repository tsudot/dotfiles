Vim�UnDo� �C�0����O�0�Wi�wIGa�u�FOg  8              incoming_pricing = IncomingNumberRate.objects.exclude(tier__tier='00').exclude(voice_rate=Decimal('0.00000')).values('tier__country', 'voice_unit', 'rental_rate').filter(product_plan=product_plan_standard).order_by('tier__country').annotate(Min('voice_rate'),  �   S                       UGQB    _�                     �   S    ����                                                                                                                                                                                                                                                                                                                                                             UGQA    �  �  �  8                 incoming_pricing = IncomingNumberRate.objects.exclude(tier__tier='00').exclude(voice_rate=Decimal('0.00000')).values('tier__country', 'voice_unit', 'rental_rate').filter(product_plan=product_plan_standard).order_by('tier__country').annotate(Min('voice_rate'),5��